package usecase_test

import (
	"strconv"
	"testing"

	"github.com/Gustibimo/fave/api/model"
	"github.com/Gustibimo/fave/api/repository/mocks"
	ucase "github.com/Gustibimo/fave/api/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockMerchantRepo := new(mocks.MerchantRepository)
	mockMerchant := &model.Merchants{
		ID:      1,
		Name:    "De hoek",
		Address: "Lawickse Allee 5a",
	}

	mockListMerchant := make([]*model.Merchants, 0)
	mockListMerchant = append(mockListMerchant, mockMerchant)
	mockMerchantRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(mockListMerchant, nil)

	mockMerchantRepo.On("GetByID", mock.AnythingOfType("int")).Return(mockMerchant, nil)
	u := ucase.NewMerchantImpl(mockMerchantRepo)
	num := int(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)
	cursorExpected := strconv.Itoa(int(mockMerchant.ID))
	assert.Equal(t, cursorExpected, nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListMerchant))

	mockMerchantRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int"))

}
