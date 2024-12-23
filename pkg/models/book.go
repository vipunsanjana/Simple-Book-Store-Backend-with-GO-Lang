package models

import (
	"github.com/jinzhu/gorm"
	"github.com/vipunsanjana/book-store/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("ID = ?", Id).Delete(book)
	return book
}
