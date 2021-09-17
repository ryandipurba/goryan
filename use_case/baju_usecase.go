package usecase

import (
	"baju/entity"
	"baju/repository"
)

type BajuUsecase interface {
	Addbaju(id int, name string, harga float32) bool
	GetList() []entity.Baju
	DeleteBaju(id int) bool
	BajuLimit(limit int) []entity.Baju
}

type bajuUsecase struct {
	br repository.BajuRepository
}

func NewBajuUseCase(br repository.BajuRepository) BajuUsecase {
	return &bajuUsecase{
		br: br,
	}
}

func (b *bajuUsecase) Addbaju(id int, name string, harga float32) bool {
	p := entity.Baju{
		Id:    id,
		Name:  name,
		Harga: harga,
	}

	b.br.AddBaju(p)
	return true

}

func (b *bajuUsecase) GetList() []entity.Baju {
	return b.br.GetList()
}

func (b *bajuUsecase) DeleteBaju(id int) bool {
	b.br.DeleteBaju(id)
	println(id)
	return true
}

func (b *bajuUsecase) BajuLimit(limit int) []entity.Baju {
	return b.br.BajuLimit(limit)
}
