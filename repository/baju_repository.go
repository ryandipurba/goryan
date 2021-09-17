package repository

import "baju/entity"

type BajuRepository interface {
	AddBaju(payload entity.Baju) bool
	GetList() []entity.Baju
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
