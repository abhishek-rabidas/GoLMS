package service

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"gomvc/model"
	"gomvc/views"
	"os"
)

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

// GetBooks TODO: pagination
func (b *BookService) GetBooks(userType string) *[]views.BookResponse {
	var books []views.BookResponse

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

		book := views.BookResponse{
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

			book := views.BookResponse{
				BookName:        record[0],
				Author:          record[1],
				PublicationYear: record[2],
			}

			books = append(books, book)
		}
	}

	return &books
}
