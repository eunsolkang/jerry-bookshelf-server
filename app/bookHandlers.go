package app

import (
	"encoding/json"
	"net/http"

	"github.com/book/service"
)

type BookHandlers struct {
	service service.BookService
}

func (bh *BookHandlers) getAllBook(writer http.ResponseWriter, request *http.Request) {

	books, err := bh.service.GetAllBook()

	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(err)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(books)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	
	w.WriteHeader(code)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
