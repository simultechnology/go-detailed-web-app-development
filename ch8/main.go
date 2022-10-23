package main

import (
	"errors"
	"fmt"
)

type AuthorID string

func (i *AuthorID) Valid() bool {
	return false
}

type Author struct {
	ID   AuthorID
	name string
}

func (a *Author) Name() string {
	return a.name
}

type Book struct {
	Title    string
	AuthorID AuthorID
}

func (b *Book) Name() string {
	return b.Title
}

func main() {
	fmt.Println("ch8 starts!!")
	book := &Book{
		"Norwegian wood",
		"123",
	}
	_, err := GetAuthorName(book)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func GetAuthor(id AuthorID) (*Author, error) {
	if !id.Valid() {
		return nil, errors.New("GetAuthor: id is invalid")
	}
	return &Author{
		ID: id,
	}, nil
}

func GetAuthorName(b *Book) (string, error) {
	author, err := GetAuthor(b.AuthorID)
	if err != nil {
		return "", fmt.Errorf("GetAuthorName: %v", err)
	}
	return author.Name(), nil
}
