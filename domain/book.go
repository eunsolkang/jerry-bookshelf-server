package domain

import (
	"github.com/book/dto"
)

type Book struct {
	Id        string
	Name      string
	Read_date string
	Report    string
	Rating    float64
	Img_url   string
}

func (b Book) ToDto() dto.BookResponse {
	return dto.BookResponse{
		Id:			b.Id,
		Name: 		b.Name,
		Img_url:    b.Img_url,
		Read_date:  b.Read_date,
		Report:     b.Report,
		Rating:     b.Rating,
	}
} 

type BookRepository interface {
	FindAll() ([]Book, error)
} 
