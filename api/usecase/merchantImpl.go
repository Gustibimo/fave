package usecase

import (
	"strconv"

	"github.com/Gustibimo/fave/api/model"
	"github.com/Gustibimo/fave/api/repository"
	"github.com/labstack/gommon/log"
)

type merchantImpl struct {
	merchantRepos repository.MerchantRepository
}

type MerchantImpl interface {
	Fetch(cursor string, num int64) ([]*model.Merchants, string, error)
	GetByID(id int64) (*model.Merchants, error)
	// GetByName(name string) (*model.Merchants, error)
	Update(merchant *model.Merchants) (*model.Merchants, error)
	Store(s *model.Merchants) (*model.Merchants, error)
	Delete(id int64) (bool, error)
}

func NewMerchantImpl(m repository.MerchantRepository) MerchantImpl {
	return &merchantImpl{
		merchantRepos: m,
	}
}

func (m *merchantImpl) Delete(id int64) (bool, error) {
	existedMerchant, _ := m.merchantRepos.GetByID(id)
	log.Info("masuk sini")
	if existedMerchant == nil {
		log.Info("Masuk Sini2")
		return false, model.NOT_FOUND_ERROR
	}
	log.Info("Masuk Sini3")

	return m.merchantRepos.Delete(id)
}

func (m *merchantImpl) Fetch(cursor string, num int64) ([]*model.Merchants, string, error) {
	if num == 0 {
		num = 10
	}

	listMerchant, err := m.merchantRepos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""

	if size := len(listMerchant); size == int(num) {
		lastId := listMerchant[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listMerchant, nextCursor, nil
}

func (m *merchantImpl) GetByID(id int64) (*model.Merchants, error) {

	res, err := m.merchantRepos.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *merchantImpl) Update(mr *model.Merchants) (*model.Merchants, error) {
	return m.merchantRepos.Update(mr)
}

func (m *merchantImpl) Store(s *model.Merchants) (*model.Merchants, error) {

	existedMerchant, _ := m.GetByID(s.ID)
	if existedMerchant != nil {
		return nil, model.CONFLIT_ERROR
	}

	id, err := m.merchantRepos.Store(s)
	if err != nil {
		return nil, err
	}
	s.ID = id

	return s, nil
}
