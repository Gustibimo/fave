package repository

import (
	model "github.com/Gustibimo/fave/api/model"
)

type MerchantRepository interface {
	Fetch(cursor string, num int64) ([]*model.Merchants, error)
	GetByID(id int64) (*model.Merchants, error)
	// GetByName(name string) (*model.Merchants, error)
	Update(merchant *model.Merchants) (*model.Merchants, error)
	Store(s *model.Merchants) (int64, error)
	Delete(id int64) (bool, error)
}
