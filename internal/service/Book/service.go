package book

import (
	"clear-arch/internal/entity"
	"clear-arch/internal/logger"
	"context"
)

type BookRepo interface {
	StoreBooks(context.Context, []entity.Book) error
	GetAllBooks(context.Context) ([]entity.Book, error)
}

type BookService struct {
	repo BookRepo
	log  logger.Logger
}

func New(r BookRepo, l logger.Logger) *BookService {
	return &BookService{
		repo: r,
		log:  l,
	}
}
