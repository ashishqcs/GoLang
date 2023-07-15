package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	mocks "internet.provider/mocks"
)

func TestShouldSaveInternetBills(t *testing.T) {
	mockClient := new(mocks.Client)
	mockRepo := new(mocks.Repository)

	ErrClientException := errors.New("Test Error")

	mockClient.On("GetBillById", mock.Anything).Return(nil, ErrClientException)
	mockClient.On("GetBillById", 1).Return("id,name,address,plan-name,date,amount\n1900,Shilpa,Kolkata,SuperSaver,01-03-2022,86", nil)

	mockRepo.On("SaveBills", mock.AnythingOfType("[]postgres.BillEntity")).Return(nil)

	mockRepo.On("GetBillsByName", mock.AnythingOfType("string")).Return(nil, nil)

	service := NewInternetProviderService(mockClient, mockRepo)
	service.GetInternetBills(0, 50)

	mockRepo.AssertCalled(t, "SaveBills", mock.AnythingOfType("[]postgres.BillEntity"))
	mockRepo.TestData().Copy().Value()
}

// func TestShouldSaveInternetBills2(t *testing.T) {

// 	service := NewInternetProviderService(client.NewInternetProviderClient(), postgres.NewBillRepository())
// 	service.GetInternetBills(1900, 2050)
// }
