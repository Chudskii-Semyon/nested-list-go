package domain

import (
	"github.com/jinzhu/gorm"
)

type List struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Items []Item
}

type ListRepository interface {
	Fetch() ([]List, error)
	GetByID(listID uint) (List, error)
	Create(list List) error
	Delete(listID uint) error
}

type ListUsecase interface {
	Fetch() ([]List, error)
	GetByID(listID uint) (List, error)
	Create(list List) error
	Delete(listID uint) error
}
