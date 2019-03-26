package cielo_api30_test

import (
	"cielo_api30"
	"testing"
)

var (
	merchant = cielo_api30.NewMerchant("","")
	env = cielo_api30.SandboxEnvironment

	ecommerce  = cielo_api30.NewEcommerce(merchant, env)
)


func TestInstance(t *testing.T) {

}

func TestCreateSale(t *testing.T) {

	var sale cielo_api30.Sale = cielo_api30.Sale{}

	ecommerce.CreateSale(sale)

}