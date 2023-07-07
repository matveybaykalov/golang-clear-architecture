package book

import (
	"clear-arch/internal/entity"
	"context"
)

func (b *BookService) GetBooks(ctx context.Context) ([]entity.Book, error) {
	return b.repo.GetAllBooks(ctx)
}
