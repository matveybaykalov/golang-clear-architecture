package usecase

import (
	"clear-arch/internal/entity"
	"context"
	"fmt"
)

type formatParser interface {
	GetBooks(context.Context) ([]entity.Book, error)
}

func (uc *UseCase) DownloadLibWithType(ctx context.Context, t string) error {
	var parser formatParser

	switch t {
	case "bmstu":
		parser = uc.bmstuParser
	case "msu":
		parser = uc.msuParser
	case "hse":
		parser = uc.hseParser
	default:
		return fmt.Errorf("unknown type: %s", t)
	}

	d, err := parser.GetBooks(ctx)

	if err != nil {
		return err
	}

	return uc.bookService.StoreBooks(ctx, d)
}
