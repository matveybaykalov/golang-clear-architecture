package bmsturepo

import (
	"clear-arch/internal/entity"
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

func (br *BmstuRepo) getBytes(ctx context.Context) ([]byte, error) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	br.log.Debug(fmt.Sprintf("project root : %s", basepath))

	path := filepath.Join(basepath, "../../../test/yml/example1.yml")

	br.log.Info("Open file test/yml/example1.yml")

	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (br *BmstuRepo) GetBooks(ctx context.Context) ([]entity.Book, error) {
	d, err := br.getBytes(ctx)
	if err != nil {
		return nil, err
	}

	br.log.Info("Data was recived")

	ymlBooks := BmstuFormat{}

	err = yaml.Unmarshal(d, &ymlBooks)
	if err != nil {
		return nil, err
	}

	br.log.Info("data was marshaled")

	res := []entity.Book{}

	br.log.Info("start convert entity")

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

	br.log.Info("stop convert entity")

	return res, nil
}
