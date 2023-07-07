package usecase

import "context"

func (uc *UseCase) DownloadFilterBmstuLib(ctx context.Context) error {
	d, err := uc.bmstuParser.GetBooks(ctx)

	if err != nil {
		return err
	}

	d, err = uc.bookService.FilterBooks(ctx, d)

	if err != nil {
		return err
	}

	return uc.bookService.StoreBooks(ctx, d)
}
