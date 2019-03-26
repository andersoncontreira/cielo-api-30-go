package cielo_api30

import (
	"bytes"
	"net/http"
)

//type Request interface {
//	execute() string
//}

type Request struct {
	url string
	environment Environment
}

//type SaleRequest struct {
//	url string
//	environment Environment
//}

func NewSaleRequest(ecommerce * Ecommerce) Request {

	env := ecommerce.getEnvironment()
	url := env.getApiUrl() + "/1/sales"

	request := Request{url:url, environment:env}

	return request
}


func (request * Request) execute (sale Sale) string {

	http.NewRequest(http.MethodPost, request.url, bytes.NewBuffer(sale.ToBuffer()))

	 return ""
}