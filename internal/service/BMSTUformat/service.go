package bmstuformat

import (
	"clear-arch/internal/entity"
	"clear-arch/internal/interfaces"
	"context"
)

type BmstuFormatRepo interface {
	GetBooks(context.Context) ([]entity.Book, error)
}

type BmstuFormatService struct {
	repo BmstuFormatRepo
	log  interfaces.Logger
}

func New(r BmstuFormatRepo, l interfaces.Logger) *BmstuFormatService {
	return &BmstuFormatService{
		repo: r,
		log:  l,
	}
}
