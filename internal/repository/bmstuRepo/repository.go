package bmsturepo

import "clear-arch/internal/logger"

type BmstuRepo struct {
	log logger.Logger
}

func New(l logger.Logger) *BmstuRepo {
	return &BmstuRepo{
		log: l,
	}
}
