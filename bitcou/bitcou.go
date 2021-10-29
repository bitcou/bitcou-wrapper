package bitcou

type Bitcou struct {
	apiKey string
	URL string
	dev bool
}

func NewBitcou(apiKey string, dev bool) * Bitcou {
	b := new(Bitcou)
	b.apiKey = apiKey
	if dev {
		b.URL = "https://sandbox-bitcou.kindynos.com/"
	} else {
		b.URL = "https://api-bitcou.kindynos.com/"
	}
	return b
}

func (* Bitcou) GetBalance() (interface{}, error) {
	return nil, nil
}
