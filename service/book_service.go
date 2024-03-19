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

	_, err = file.WriteString(fmt.Sprintf("\n%s,%s,%s", request.BookName, request.Author, request.PublicationYear))
	if err != nil {
		log.Error("Unable to add new book")
		return err
	}

	return nil
}

func (b *BookService) DeleteBook(bookName string) error {

	if strings.TrimSpace(bookName) == "" {
		return exception.New("Book Name is missing")
	}

	file, err := os.Open("resources/regularUser.csv")

	if err != nil {
		log.Error(err)
		return err
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Error(err)
		return nil
	}

	var books []views.BookDTO

	for i, record := range records {
		if i == 0 {
			continue
		}

		book := views.BookDTO{
			BookName:        fmt.Sprintf("\"%s\"", record[0]),
			Author:          record[1],
			PublicationYear: record[2],
		}

		books = append(books, book)
	}

	file.Close()

	err = os.Truncate("resources/regularUser.csv", 0)

	if err != nil {
		return exception.New("Something went wrong while deleting book")
	}

	file, err = os.OpenFile("resources/regularUser.csv", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Error(err)
		return err
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s,%s,%s", "Book Name", "Author", "Publication Year"))
	if err != nil {
		log.Error("Unable to delete book")
		return err
	}

	for _, record := range books {

		if strings.ToLower(strings.Trim(record.BookName, "\"")) == strings.ToLower(bookName) {
			continue
		} else {
			_, err = file.WriteString(fmt.Sprintf("\n%s,%s,%s", record.BookName, record.Author, record.PublicationYear))
			if err != nil {
				log.Error("Unable to delete book")
				return err
			}
		}
	}

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
