package bmstuformat

import (
	"clear-arch/internal/logger"
	"context"
)

type BmstuFormatRepo interface {
	GetBytes(context.Context) ([]byte, error)
}

type BmstuFormatService struct {
	repo BmstuFormatRepo
	log  logger.Logger
}

func New(r BmstuFormatRepo, l logger.Logger) *BmstuFormatService {
	return &BmstuFormatService{
		repo: r,
		log:  l,
	}
}
