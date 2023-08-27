package book

import (
	"clear-arch/internal/entity"
	"clear-arch/internal/interfaces"
	"context"
)

type BookRepo interface {
	StoreBooks(context.Context, []entity.Book) error
	GetAllBooks(context.Context) ([]entity.Book, error)
}

type BookService struct {
	repo BookRepo
	log  interfaces.Logger
}

func New(r BookRepo, l interfaces.Logger) *BookService {
	return &BookService{
		repo: r,
		log:  l,
	}
}
