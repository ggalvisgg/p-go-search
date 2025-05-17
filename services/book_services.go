package services

import (
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/repositories"
)

type BookServiceInterface interface {
    GetBooks() ([]models.Book, error)
}

type BookService struct {
    repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
    return &BookService{repo}
}

func (s *BookService) GetBooks() ([]models.Book, error) {
    return s.repo.GetAllBooks()
}
