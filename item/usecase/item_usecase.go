package usecase

import "awesomeProject/domain"

type itemUsecase struct {
	itemRepo domain.ItemRepository
}

func NewItemUsecase(l domain.ItemRepository) domain.ItemUsecase {
	return &itemUsecase{l}
}

func (i itemUsecase) GetAll(listID uint) (res []domain.Item, err error) {
	return i.itemRepo.GetAll(listID)
}

func (i itemUsecase) GetByID(listID, itemID uint) (res domain.Item, err error) {
	return i.itemRepo.GetByID(listID, itemID)
}

func (i itemUsecase) Create(listID uint, item domain.Item) error {
	return i.itemRepo.Create(listID, item)
}

func (i itemUsecase) Update(item domain.Item) error {
	return i.itemRepo.Update(item)
}

func (i itemUsecase) Delete(listID, itemID uint) error {
	return i.itemRepo.Delete(listID, itemID)
}
