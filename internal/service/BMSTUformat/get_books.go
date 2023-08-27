package bmstuformat

import (
	"clear-arch/internal/entity"
	"context"
)

func (bmstu *BmstuFormatService) GetBooks(ctx context.Context) ([]entity.Book, error) {
	return bmstu.repo.GetBooks(ctx)
}
