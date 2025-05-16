package controllers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "example.com/go-mongo-app/models"
)

// ---------------------- MOCK DEL SERVICIO ----------------------


type MockBookService struct {
    mock.Mock
}

func (m *MockBookService) GetBooks() ([]models.Book, error) {
    args := m.Called()
    return args.Get(0).([]models.Book), args.Error(1)
}
// ---------------------- TESTS ----------------------


func TestGetBooks_Success(t *testing.T) {
    mockService := new(MockBookService)
    controller := NewBookController(mockService)

    books := []models.Book{
        {Title: "Book 1", Author: "Author 1", ISBN: "1111111111"},
        {Title: "Book 2", Author: "Author 2", ISBN: "2222222222"},
    }

    mockService.On("GetBooks").Return(books, nil)

    req := httptest.NewRequest("GET", "/books", nil)
    resp := httptest.NewRecorder()

    controller.GetBooks(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    mockService.AssertExpectations(t)
}