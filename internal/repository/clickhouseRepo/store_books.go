package clickhouserepo

import (
	"clear-arch/internal/entity"
	"context"
	"math/rand"

	"dev.solidwall.io/solidlab/NGPlatform/ost/go-clickhouse.git/ch"
)

func getModels(ctx context.Context, d []entity.Book) ([]book, []author) {
	books := []book{}
	authors := []author{}

	for _, e := range d {
		bId := rand.Intn(10000)
		aId := rand.Intn(10000)

		b := book{
			Id:        bId,
			AuthorId:  aId,
			Name:      e.Name,
			PageCount: e.PackageCount,
		}

		a := author{
			Id:      aId,
			Name:    e.Author.Name,
			Surname: e.Author.Surname,
		}

		books = append(books, b)
		authors = append(authors, a)
	}

	return books, authors
}

func (chr *ClickHouseRepo) StoreBooks(ctx context.Context, d []entity.Book) error {
	chr.log.Info("start upload to clickhouse")

	booksTableName := "books_distr"
	authorsTableName := "authors_distr"

	b, a := getModels(ctx, d)

	chr.log.Info("got models")

	_, err := chr.con.NewInsert().Model(&b).
		ModelTableExpr("?", ch.Safe(booksTableName)).Exec(context.Background())

	if err != nil {
		chr.log.Info("error was occurred")
		return err
	}

	chr.log.Info("save books to clickhouse")

	_, err = chr.con.NewInsert().Model(&a).
		ModelTableExpr("?", ch.Safe(authorsTableName)).Exec(context.Background())

	return err
}
