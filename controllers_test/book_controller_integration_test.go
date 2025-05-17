package controllers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "example.com/go-mongo-app/controllers"
    "example.com/go-mongo-app/repositories"
    "example.com/go-mongo-app/services"
    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
    repo := repositories.NewBookRepository()
    service := services.NewBookService(repo)
    controller := controllers.NewBookController(service)

    r := mux.NewRouter()
    r.HandleFunc("/books", controller.GetBooks).Methods("GET")

    return r
}

func TestBookCRUDIntegration(t *testing.T) {
    router := setupRouter()

    reqGetAll, _ := http.NewRequest("GET", "/books", nil)
    rrGetAll := httptest.NewRecorder()
    router.ServeHTTP(rrGetAll, reqGetAll)

    assert.Equal(t, http.StatusOK, rrGetAll.Code)
}
