package book

import (
	"clear-arch/internal/entity"
	"context"
	"errors"
)

func (b *BookService) FilterBooks(context.Context, []entity.Book) ([]entity.Book, error) {
	// TODO: implement

	return nil, errors.New("Implement FilterBooks method in BookService")
}
