package clickhouserepo

import (
	"clear-arch/internal/interfaces"

	"dev.solidwall.io/solidlab/NGPlatform/ost/go-clickhouse.git/ch"
)

type ClickHouseRepo struct {
	con *ch.DB
	log interfaces.Logger
}

func New(c *ch.DB, l interfaces.Logger) *ClickHouseRepo {
	return &ClickHouseRepo{
		con: c,
		log: l,
	}
}
