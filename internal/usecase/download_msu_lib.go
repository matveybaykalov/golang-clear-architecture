package usecase

import (
	"context"
)

func (uc *UseCase) DownloadMsuLib(ctx context.Context) error {
	d, err := uc.msuParser.GetBooks(ctx)

	if err != nil {
		return err
	}

	uc.log.Info("get books entities")

	return uc.bookService.StoreBooks(ctx, d)
}
