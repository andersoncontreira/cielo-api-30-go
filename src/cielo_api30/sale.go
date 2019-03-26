package cielo_api30

import "encoding/json"

type Sale struct {
	MerchantOrderID string    `json:",omitempty"`
	Customer        *Customer `json:",omitempty"`
	Payment         *Payment  `json:",omitempty"`
}

func (sale Sale) ToBuffer() []byte {

	body, _ := json.Marshal(sale)

	return body
}

func (sale Sale) ToJson() string {
	body, _ := json.Marshal(sale)

	return string(body)
}

//fromJson()
//populate