package usecase

import (
	"baju/entity"
	"baju/repository"
)

type BajuUsecase interface {
	// AddBaju(id int, name string) bool
	GetList() []entity.Baju
}

type bajuUsecase struct {
	br repository.BajuRepository
}

func NewBajuUseCase(br repository.BajuRepository) BajuUsecase {
	return &bajuUsecase{
		br: repository.NewBajuRepository(),
	}
}

// func (b *bajuUsecase) AddBaju(id int, name string) bool {
// 	p := entity.Baju{
// 		Id:   id,
// 		Name: name,
// 	}

// 	b.AddBaju(p.Id, p.Name)
// 	return true
// }

func (b *bajuUsecase) GetList() []entity.Baju {
	return b.GetList()
}
