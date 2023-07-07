package usecase

import "context"

func (uc *UseCase) DownloadBmstuLib(ctx context.Context) error {
	d, err := uc.bmstuParser.GetBooks(ctx)

	if err != nil {
		return err
	}

	uc.log.Info("get books entities")

	return uc.bookService.StoreBooks(ctx, d)
}
