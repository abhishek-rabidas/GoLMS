package service

import "gomvc/views"

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

func (b *BookService) GetBooks(userType string) *[]views.BookResponse {
	return nil
}
