package domain

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name     string `gorm:"not null"`
	ListID   uint   `gorm:"type:int REFERENCES lists(id) ON DELETE CASCADE"`
	Position uint   `gorm:"not null"`
}

type ItemRepository interface {
	GetAll(listID uint) (res []Item, err error)
	GetByID(listID, itemID uint) (Item, error)
	Create(listID uint, item Item) error
	Update(item Item) error
	Delete(listID, itemID uint) error
}

type ItemUsecase interface {
	GetAll(listID uint) (res []Item, err error)
	GetByID(listID, itemID uint) (res Item, err error)
	Create(listID uint, item Item) error
	Update(item Item) error
	Delete(listID, itemID uint) error
}
