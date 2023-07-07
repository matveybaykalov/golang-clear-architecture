package msuformat

import (
	"clear-arch/internal/entity"
	"context"
	"encoding/xml"
)

func (msu *MsuFormatService) GetBooks(ctx context.Context) ([]entity.Book, error) {
	d, err := msu.repo.GetBytes(ctx)
	if err != nil {
		msu.log.Error()
		return nil, err
	}

	xmlBooks := MsuFormat{}

	err = xml.Unmarshal(d, &xmlBooks)
	if err != nil {
		return nil, err
	}

	res := []entity.Book{}

	authors := make(map[int]entity.Author, 0)

	for _, a := range xmlBooks.Authors.Author {
		authors[a.Id] = struct {
			Name    string
			Surname string
		}{
			Name:    a.Name,
			Surname: a.Surname,
		}
	}

	for _, b := range xmlBooks.Books.Book {
		res = append(res, entity.Book{
			Name:         b.Name,
			PackageCount: b.PageCount,
			Author:       authors[b.AuthorId],
		})
	}

	return res, nil
}
