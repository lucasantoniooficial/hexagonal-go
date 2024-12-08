package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lucasantoniooficial/hexagonal-go/adapters/cli"
	mock_application "github.com/lucasantoniooficial/hexagonal-go/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	price := 20.0
	productStatus := "disabled"
	productId := "1"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(ctrl)
	productServiceMock.EXPECT().Create(productName, price).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", productId, productName, price, productStatus)

	result, err := cli.Run(productServiceMock, "create", "", productName, price)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s has been enabled", productId)

	result, err = cli.Run(productServiceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s has been disabled", productId)

	result, err = cli.Run(productServiceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has the price %f and status %s", productId, productName, price, productStatus)

	result, err = cli.Run(productServiceMock, "", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
