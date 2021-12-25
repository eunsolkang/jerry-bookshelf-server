package service

import (
	"github.com/book/domain"
	"github.com/book/dto"
)

type BookService interface {
	GetAllBook() ([]dto.BookResponse, error)
}

type DefaultBookService struct {
	repo domain.BookRepository
}

func (s DefaultBookService) GetAllBook() ([]dto.BookResponse, error) {
	bs, err := s.repo.FindAll()
	response := make([]dto.BookResponse, 0)

	if err != nil {
		return nil, err
	}

	for _, b := range bs {
		response = append(response, b.ToDto())
	}

	return response, nil
}

func NewBookService(repository domain.BookRepository) DefaultBookService {
	return DefaultBookService{repository}
}
