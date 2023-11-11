package application_test

import (
	"github.com/sergiodanilojr/hexagonal-architecture-go/application"
	"github.com/stretchr/testify/require"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestProduct_Enable(t *testing.T){
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

func TestProduct_Disable(t *testing.T){
	product := application.Product{}
	product.Name = "Caixa de Papelão"
	product.Status = application.ENABLED
	product.Price = 0
	
	err := product.Disable()

	require.Nil(t, err)

	product.Price = 20
	err = product.Disable()
	
	require.Equal(t, "The price must be zero in order in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T){
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Fôrma de Bolo Inglês"
	product.Status = application.ENABLED
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
