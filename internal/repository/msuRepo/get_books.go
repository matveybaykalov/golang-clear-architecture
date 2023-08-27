package msurepo

import (
	"clear-arch/internal/entity"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func (mr *MsuRepo) getBytes(ctx context.Context) ([]byte, error) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	mr.log.Debug(fmt.Sprintf("project root : %s", basepath))

	path := filepath.Join(basepath, "../../../test/xml/example1.xml")

	mr.log.Info("Open file test/yml/example1.xml")

	d, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (mr *MsuRepo) GetBooks(ctx context.Context) ([]entity.Book, error) {
	d, err := mr.getBytes(ctx)
	if err != nil {
		mr.log.Error()
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
