package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Price     float64 `json:"price"`
	Publisher string  `json:"publisher"`
	Pages     int     `json:"pages"`
}

func main() {
	b1 := Book{
		Title:     "Atomic Habits",
		Author:    "James Clear",
		Price:     200,
		Publisher: "Avery",
		Pages:     300,
	}

	jsondata, err := json.MarshalIndent(b1, "", " ")
	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}
	fmt.Println("Struct to json")
	fmt.Println(string(jsondata))

	var b2 Book
	err = json.Unmarshal(jsondata, &b2)
	if err != nil {
		fmt.Println("Error Unmarshalling:", err)
		return
	}
	fmt.Println("\nJSON to struct")
	fmt.Printf("Title: %s\nAuthor: %s\nPrice: %.2f\nPublisher:%s\nPages:%d\n", b2.Title, b2.Author, b2.Price, b2.Publisher, b2.Pages)

}
