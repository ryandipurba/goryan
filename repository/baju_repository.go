package repository

import (
	"baju/entity"
)

type BajuRepository interface {
	AddBaju(payload entity.Baju) bool
	GetList() []entity.Baju
	DeleteBaju(id int) bool
	BajuLimit(limit int) []entity.Baju
}

type bajuRepository struct {
	baju []entity.Baju
}

func NewBajuRepository() BajuRepository {
	return &bajuRepository{}
}

func (b *bajuRepository) AddBaju(payload entity.Baju) bool {
	b.baju = append(b.baju, payload)
	return true
}

func (b *bajuRepository) GetList() []entity.Baju {
	return b.baju
}

func (b *bajuRepository) DeleteBaju(id int) bool {
	for i, key := range b.baju {
		if key.Id == id {
			b.baju = append(b.baju[:i], b.baju[i+1:]...)
		}
	}
	return true
}

func (b *bajuRepository) BajuLimit(limit int) []entity.Baju {
	b.baju = b.baju[0:limit]
	return b.baju
}
