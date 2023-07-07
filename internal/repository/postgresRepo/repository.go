package postgresrepo

import (
	"clear-arch/internal/logger"

	"github.com/uptrace/bun"
)

type PostgresRepo struct {
	con *bun.DB
	log logger.Logger
}

func New(c *bun.DB, l logger.Logger) *PostgresRepo {
	return &PostgresRepo{
		con: c,
		log: l,
	}
}
