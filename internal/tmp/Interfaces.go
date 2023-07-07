package usecase

import (
	"clear-arch/internal/entity"
	"context"
)

type (
	Books interface {
		GetBooks(context.Context, string) ([]entity.Book, error)
	}

	BooksRepo interface {
		GetRawData(context.Context, string) ([]byte, error)
		StoreBook(context.Context, entity.Book) error
		StoreBooks(context.Context, []entity.Book) error
		GetBooksByName(context.Context, string) ([]entity.Book, error)
	}
)
