package bmsturepo

import "clear-arch/internal/interfaces"

type BmstuRepo struct {
	log interfaces.Logger
}

func New(l interfaces.Logger) *BmstuRepo {
	return &BmstuRepo{
		log: l,
	}
}
