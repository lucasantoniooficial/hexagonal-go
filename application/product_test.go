package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/lucasantoniooficial/hexagonal-go/application"
	uuid "github.com/satori/go.uuid"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Disable()
	require.Equal(t, "The price must be zero to disable the product", err.Error())

	product.Price = 0

	err = product.Disable()
	require.Nil(t, err)
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()

	require.Equal(t, "The price must be greater than zero", err.Error())
}

func TestPrdouct_GetID (t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()

	require.Equal(t, product.ID, product.GetID())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"

	require.Equal(t, product.Name, product.GetName())
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	require.Equal(t, product.Price, product.GetPrice())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED

	require.Equal(t, product.Status, product.GetStatus())
}