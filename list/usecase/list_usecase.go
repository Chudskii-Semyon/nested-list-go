package usecase

import (
	"awesomeProject/domain"
)

type listUsecase struct {
	listRepo domain.ListRepository
}

func NewListUsecase(l domain.ListRepository) domain.ListUsecase {
	return &listUsecase{l}
}

func (l *listUsecase) Fetch() (res []domain.List, err error) {
	return l.listRepo.Fetch()
}

func (l *listUsecase) GetByID(listID uint) (res domain.List, err error) {
	return l.listRepo.GetByID(listID)
}

func (l *listUsecase) Create(list domain.List) error {
	return l.listRepo.Create(list)
}

func (l *listUsecase) Delete(listID uint) error {
	return l.listRepo.Delete(listID)
}
