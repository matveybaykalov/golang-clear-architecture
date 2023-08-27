package msuformat

import (
	"clear-arch/internal/entity"
	"clear-arch/internal/interfaces"
	"context"
)

type MsuFormatRepo interface {
	GetBooks(context.Context) ([]entity.Book, error)
}

type MsuFormatService struct {
	repo MsuFormatRepo
	log  interfaces.Logger
}

func New(r MsuFormatRepo, l interfaces.Logger) *MsuFormatService {
	return &MsuFormatService{
		repo: r,
		log:  l,
	}
}
