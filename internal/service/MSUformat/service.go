package msuformat

import (
	"clear-arch/internal/logger"
	"context"
)

type MsuFormatRepo interface {
	GetBytes(context.Context) ([]byte, error)
}

type MsuFormatService struct {
	repo MsuFormatRepo
	log  logger.Logger
}

func New(r MsuFormatRepo, l logger.Logger) *MsuFormatService {
	return &MsuFormatService{
		repo: r,
		log:  l,
	}
}
