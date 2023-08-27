package msuformat

import (
	"clear-arch/internal/entity"
	"context"
)

func (msu *MsuFormatService) GetBooks(ctx context.Context) ([]entity.Book, error) {
	return msu.repo.GetBooks(ctx)
}
