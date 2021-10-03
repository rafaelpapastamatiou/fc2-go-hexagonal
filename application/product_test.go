package application_test

import (
	"testing"

	"github.com/rafaelpapastamatiou/fc2-go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
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

	require.Equal(
		t,
		application.ERROR_ENABLING_PRODUCT,
		err.Error(),
	)
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()

	require.Equal(
		t,
		application.ERROR_DISABLING_PRODUCT,
		err.Error(),
	)
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
	require.Equal(t, application.ERROR_STATUS_INVALID, err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, application.ERROR_PRICE_INVALID, err.Error())

	product.Price = 10
	_, err = product.IsValid()
	require.Nil(t, err)
}
