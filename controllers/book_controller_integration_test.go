package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-mongo-app/models"
	"example.com/go-mongo-app/repositories"
	"example.com/go-mongo-app/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)
	controller := NewBookController(service)

	r := mux.NewRouter()
	r.HandleFunc("/books", controller.GetBooks).Methods("GET")

	return r
}

func TestBookCRUDIntegration(t *testing.T) {
	router := setupRouter()

	// 1. Get all books
	reqGetAll, _ := http.NewRequest("GET", "/books", nil)
	rrGetAll := httptest.NewRecorder()
	router.ServeHTTP(rrGetAll, reqGetAll)
	assert.Equal(t, http.StatusOK, rrGetAll.Code)

}
