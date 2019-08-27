package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Page holds metadata about a pages.
type Page struct {
	gorm.Model
	NumPage uint64
	Text    string
}

// Book holds metadata about a book.
type Book struct {
	gorm.Model
	Author    string
	Publisher string
	PageCount int64
	Page      []Page
}

// Shelf holds metadata about a shelf.
type Shelf struct {
	gorm.Model
	NumShelf int
	Title    string
	Book     []Book
}

func getAllShelfs(db *gorm.DB) ([]Shelf, error) {
	//rows, err := db.Find()
	shelfs := []Shelf{}
	var err error
	//var db *gorm.DB
	db.Find(&shelfs)
	//a.Rows()
	//fmt.Fprintf(w, "Shelfs: %v\n", a)

	if err != nil {
		return nil, err
	}

	//defer rows.Close()

	//shelfs := []Shelf{}

	//for rows.Rows() {
	//	var p product
	//	if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
	//		return nil, err
	//	}
	//	products = append(products, p)
	//}

	return shelfs, nil
}

func (s *Shelf) getBooks(db *gorm.DB) error {
	db.First(s.NumShelf)
	return nil
}
