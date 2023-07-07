package book

import (
	"clear-arch/internal/entity"
	"context"
)

func (b *BookService) StoreBooks(ctx context.Context, d []entity.Book) error {
	return b.repo.StoreBooks(ctx, d)
}
