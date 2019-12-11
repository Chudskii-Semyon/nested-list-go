package postgres

import (
	"awesomeProject/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type postgresListRepository struct {
	Conn *gorm.DB
}

func NewPostgresListRepository(Conn *gorm.DB) domain.ListRepository {
	return &postgresListRepository{Conn}
}

func (r *postgresListRepository) Fetch() (res []domain.List, err error) {
	if err := r.Conn.Find(&res).Error; err != nil {
		return nil, err
	}

	return
}

func (r postgresListRepository) GetByID(listID uint) (res domain.List, err error) {
	if err := r.Conn.First(&res, listID).Error; err != nil {
		return domain.List{}, err
	}

	return
}

func (r *postgresListRepository) Create(list domain.List) (err error) {
	return r.Conn.Create(&list).Error
}

func (r *postgresListRepository) Delete(listID uint) error {
	var list domain.List
	err := r.Conn.First(&list, listID).Error

	if err != nil {
		fmt.Println("Cant find list to delete. Error => ", err)
	}

	return r.Conn.Unscoped().Delete(&list).Error
}
