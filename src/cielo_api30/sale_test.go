package cielo_api30_test

import (
	"fmt"
	"cielo_api30"
	"testing"
)

func TestToBuffer(t *testing.T) {

	MerchantOrderID := "123456"
	var customer cielo_api30.Customer = cielo_api30.Customer{}
	var payment cielo_api30.Payment = cielo_api30.Payment{}

	sale := cielo_api30.Sale{
		MerchantOrderID,
		&customer,
		&payment,
	}

	fmt.Print(sale.ToJson())

}