package msurepo

import "clear-arch/internal/interfaces"

type MsuRepo struct {
	log interfaces.Logger
}

func New(l interfaces.Logger) *MsuRepo {
	return &MsuRepo{
		log: l,
	}
}
