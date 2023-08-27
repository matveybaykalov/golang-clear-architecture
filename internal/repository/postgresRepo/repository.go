package postgresrepo

import (
	"clear-arch/internal/interfaces"

	"github.com/uptrace/bun"
)

type PostgresRepo struct {
	con *bun.DB
	log interfaces.Logger
}

func New(c *bun.DB, l interfaces.Logger) *PostgresRepo {
	return &PostgresRepo{
		con: c,
		log: l,
	}
}
