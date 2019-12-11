package postgres

import (
	"awesomeProject/domain"
	"github.com/jinzhu/gorm"
)

type postgresItemRepository struct {
	Conn *gorm.DB
}

func NewPostgresItemRepository(Conn *gorm.DB) domain.ItemRepository {
	return &postgresItemRepository{Conn}
}

func (r *postgresItemRepository) GetAll(listID uint) (res []domain.Item, err error) {
	list := domain.List{
		Model: gorm.Model{ID: listID},
	}
	if err := r.Conn.Model(&list).Related(&res).Error; err != nil {
		return nil, err
	}

	return
}

func (r *postgresItemRepository) GetByID(listID, itemID uint) (res domain.Item, err error) {
	err = r.Conn.First(&res, itemID).Error
	return
}

func (r *postgresItemRepository) Create(listID uint, item domain.Item) error {
	item.ListID = listID

	return r.Conn.Create(&item).Error
}

func (r *postgresItemRepository) Update(item domain.Item) error {
	return r.Conn.Update(&item).Error
}

func (r *postgresItemRepository) Delete(listID uint, itemID uint) error {
	var item domain.List

	err := r.Conn.First(&item, itemID).Error

	if err != nil {
		return err
	}

	return r.Conn.Unscoped().Delete(&item).Error
}
