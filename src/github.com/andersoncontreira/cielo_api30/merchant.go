package cielo_api30

type Merchant struct {
	id string
	key string
}

func NewMerchant(id string, key string) Merchant {
	return Merchant{
		id: id,
		key: key,
	}
}