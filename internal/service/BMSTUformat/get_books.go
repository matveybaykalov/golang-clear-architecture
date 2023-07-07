package bmstuformat

import (
	"clear-arch/internal/entity"
	"context"

	"gopkg.in/yaml.v3"
)

func (bmstu *BmstuFormatService) GetBooks(ctx context.Context) ([]entity.Book, error) {
	d, err := bmstu.repo.GetBytes(ctx)
	if err != nil {
		return nil, err
	}

	bmstu.log.Info("Data was recived")

	ymlBooks := BmstuFormat{}

	err = yaml.Unmarshal(d, &ymlBooks)
	if err != nil {
		return nil, err
	}

	bmstu.log.Info("data was marshaled")

	res := []entity.Book{}

	bmstu.log.Info("start convert entity")

	for _, d := range ymlBooks.Books {
		res = append(res, entity.Book{
			Name:         d.Name,
			PackageCount: d.PageCount,
			Author: entity.Author{
				Name:    d.AuthorName,
				Surname: d.AuthorSurname,
			},
		})
	}

	bmstu.log.Info("stop convert entity")

	return res, nil
}
