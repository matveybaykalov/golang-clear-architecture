package clickhouserepo

import (
	"clear-arch/internal/logger"

	"dev.solidwall.io/solidlab/NGPlatform/ost/go-clickhouse.git/ch"
)

type ClickHouseRepo struct {
	con *ch.DB
	log logger.Logger
}

func New(c *ch.DB, l logger.Logger) *ClickHouseRepo {
	return &ClickHouseRepo{
		con: c,
		log: l,
	}
}
