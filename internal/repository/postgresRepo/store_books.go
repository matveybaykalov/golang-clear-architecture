package postgresrepo

import (
	"clear-arch/internal/entity"
	"context"
	"fmt"
	"math/rand"
)

func getModels(ctx context.Context, d []entity.Book) ([]book, []author) {
	books := []book{}
	authors := []author{}

	for _, e := range d {
		bId := rand.Intn(1000000)
		aId := rand.Intn(1000000)

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

func (pgr *PostgresRepo) StoreBooks(ctx context.Context, d []entity.Book) error {
	pgr.log.Info("start upload to postgres")

	booksTableName := "books"
	authorsTableName := "authors"

	b, a := getModels(ctx, d)

	pgr.log.Debug(fmt.Sprintf("books count: %d", len(b)))

	_, err := pgr.con.NewInsert().Model(&b).ModelTableExpr(booksTableName).Exec(ctx)

	if err != nil {
		pgr.log.Info("Error was occurred")
		return err
	}

	pgr.log.Info("books was uploaded")

	_, err = pgr.con.NewInsert().Model(&a).ModelTableExpr(authorsTableName).Exec(ctx)

	return err
}
