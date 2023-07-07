package clickhouserepo

import (
	"clear-arch/internal/entity"
	"context"

	"dev.solidwall.io/solidlab/NGPlatform/ost/go-clickhouse.git/ch"
)

func getEntity(books []book, authors []author) []entity.Book {
	var res []entity.Book
	authorsMap := make(map[int]entity.Author, 0)

	for _, a := range authors {
		authorsMap[a.Id] = struct {
			Name    string
			Surname string
		}{
			Name:    a.Name,
			Surname: a.Surname,
		}
	}

	for _, b := range books {
		res = append(res, entity.Book{
			Name:         b.Name,
			PackageCount: b.PageCount,
			Author:       authorsMap[b.AuthorId],
		})
	}

	return res
}

func (chr *ClickHouseRepo) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	var b []book
	var a []author

	booksTableName := "books_distr"
	authorsTableName := "authors_distr"

	err := chr.con.NewSelect().Model(&b).ModelTableExpr("?", ch.Safe(booksTableName)).Scan(ctx)

	if err != nil {
		return nil, err
	}

	err = chr.con.NewSelect().Model(&a).ModelTableExpr("?", ch.Safe(authorsTableName)).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return getEntity(b, a), nil
}
