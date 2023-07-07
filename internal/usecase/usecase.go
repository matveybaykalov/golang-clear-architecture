package usecase

import (
	"clear-arch/internal/entity"
	"clear-arch/internal/logger"
	"context"
)

type MsuFormatService interface {
	GetBooks(context.Context) ([]entity.Book, error)
}

type BmstuFormatService interface {
	GetBooks(context.Context) ([]entity.Book, error)
}

type BookService interface {
	FilterBooks(context.Context, []entity.Book) ([]entity.Book, error)
	StoreBooks(context.Context, []entity.Book) error
}

type UseCase struct {
	bmstuParser BmstuFormatService
	msuParser   MsuFormatService
	hseParser   MsuFormatService
	bookService BookService
	log         logger.Logger
}

func New(msu MsuFormatService, bmstu BmstuFormatService, hse MsuFormatService, b BookService, l logger.Logger) *UseCase {
	return &UseCase{
		bmstuParser: bmstu,
		msuParser:   msu,
		hseParser:   hse,
		bookService: b,
		log:         l,
	}
}
