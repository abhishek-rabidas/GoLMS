package service

import (
	"encoding/csv"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gomvc/exception"
	"gomvc/model"
	"gomvc/views"
	"os"
	"strings"
)

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

// GetBooks TODO: pagination
func (b *BookService) GetBooks(userType string) *[]views.BookDTO {
	var books []views.BookDTO

	file, err := os.Open("resources/regularUser.csv")

	if err != nil {
		log.Error(err)
		return nil
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Error(err)
		return nil
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		book := views.BookDTO{
			BookName:        record[0],
			Author:          record[1],
			PublicationYear: record[2],
		}

		books = append(books, book)
	}

	if userType == model.Admin {
		file, err = os.Open("resources/adminUser.csv")

		if err != nil {
			log.Error(err)
			return nil
		}

		defer file.Close()

		csvReader = csv.NewReader(file)
		records, err = csvReader.ReadAll()

		if err != nil {
			log.Error(err)
			return nil
		}

		for i, record := range records {
			if i == 0 {
				continue
			}

			book := views.BookDTO{
				BookName:        record[0],
				Author:          record[1],
				PublicationYear: record[2],
			}

			books = append(books, book)
		}
	}

	return &books
}

func (b *BookService) AddBook(request views.BookDTO) error {

	err := validateFields(&request)

	if err != nil {
		return err
	}

	file, err := os.OpenFile("resources/regularUser.csv", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Error(err)
		return err
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("\n%s,%s,%s\n", request.BookName, request.Author, request.PublicationYear))
	if err != nil {
		log.Error("Unable to add new book")
		return err
	}

	/*	csvwriter := csv.NewWriter(file)

		record := []string{request.BookName, request.Author, request.PublicationYear}

		err = csvwriter.Write(record)
		if err != nil {
			log.Error("Unable to add new book")
			return err
		}*/

	return nil

}

func validateFields(request *views.BookDTO) error {
	if strings.TrimSpace(request.BookName) == "" {
		return exception.New("Book Name is missing")
	}

	if strings.TrimSpace(request.Author) == "" {
		return exception.New("Author is missing")
	}

	if strings.TrimSpace(request.PublicationYear) == "" {
		return exception.New("Publication year is missing")
	}

	return nil
}
