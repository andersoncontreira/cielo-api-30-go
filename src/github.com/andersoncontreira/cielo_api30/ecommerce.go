package cielo_api30

type Ecommerce struct {
		merchant Merchant
		environment Environment
}


func (ecommerce * Ecommerce) CreateSale(sale Sale) {

	request := NewSaleRequest(ecommerce)
	request.execute(sale)
}

func (ecommerce * Ecommerce) getEnvironment() Environment {
	return  ecommerce.environment
}


//func NewEcommerce(merchant Merchant) * Ecommerce {
//
//	var environment Environment = SandboxEnvironment
//	return &Ecommerce{merchant:merchant, environment:environment}
//}

func NewEcommerce(merchant Merchant, environment Environment) * Ecommerce {

	return &Ecommerce{merchant:merchant, environment:environment}
}