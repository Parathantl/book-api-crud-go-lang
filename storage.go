package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const filePath = "books.json"

func loadBooks() ([]Book, error) {
	var books []Book
	file, err := os.Open(filePath)
	if err != nil {
		return books, nil
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &books)
	return books, nil
}

func saveBooks(books []Book) error {
	bytes, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, bytes, 0644)
}
