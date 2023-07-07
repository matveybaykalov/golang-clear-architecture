package filereader

import (
	"context"
	"io/ioutil"
)

type BooksParserRepo struct {
	Filename string
}

func New(fn string) *BooksParserRepo {
	return &BooksParserRepo{
		Filename: fn,
	}
}

func (pr *BooksParserRepo) GetDataFromFile(ctx context.Context) ([]byte, error) {
	d, err := ioutil.ReadFile(pr.Filename)
	if err != nil {
		return nil, err
	}

	return d, nil
}
