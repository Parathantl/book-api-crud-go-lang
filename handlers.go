package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := loadBooks()
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.BookID = uuid.New().String()
	books, _ := loadBooks()
	books = append(books, book)
	saveBooks(books)
	json.NewEncoder(w).Encode(book)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	books, _ := loadBooks()
	for _, book := range books {
		if book.BookID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updated Book
	json.NewDecoder(r.Body).Decode(&updated)

	books, _ := loadBooks()
	for i, book := range books {
		if book.BookID == params["id"] {
			updated.BookID = book.BookID
			books[i] = updated
			saveBooks(books)
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	books, _ := loadBooks()
	for i, book := range books {
		if book.BookID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			saveBooks(books)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func searchBooks(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	books, _ := loadBooks()

	type result struct {
		Book Book
		Found bool
	}
	resultChan := make(chan result)

	for _, book := range books {
		go func(b Book) {
			if strings.Contains(strings.ToLower(b.Title), query) ||
				strings.Contains(strings.ToLower(b.Description), query) {
				resultChan <- result{b, true}
			} else {
				resultChan <- result{b, false}
			}
		}(book)
	}

	var results []Book
	for range books {
		res := <-resultChan
		if res.Found {
			results = append(results, res.Book)
		}
	}
	json.NewEncoder(w).Encode(results)
}
