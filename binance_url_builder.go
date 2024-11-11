package binance_url_builder

import (
	"net/url"
	"strings"
)

type BinanceURLBuilder struct {
	url.URL
}

func (bub *BinanceURLBuilder) New(test bool) {
	bub.Scheme = string(HTTPS)
	bub.Path = string(BASE_PATH)
	if test {
		bub.Host = string(TEST)
	} else {
		bub.Host = string(PRODUCTION)
	}
}

func (bub *BinanceURLBuilder) clean() {
	bub.RawQuery = ""
}

func (bub *BinanceURLBuilder) Klines(params map[string]string) *BinanceURLBuilder {
	bub.clean()
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(KLINES),
	}, "/")
	// @todo : (Later implementation) add checks on the fields, if some field is not valid, error
	query := bub.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	bub.RawQuery = query.Encode()
	return bub
}

func (bub *BinanceURLBuilder) Order() *BinanceURLBuilder {
	bub.clean()
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(ORDER),
	}, "/")
	return bub
}

func (bub *BinanceURLBuilder) Account() *BinanceURLBuilder {
	bub.clean()
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(ACCOUNT),
	}, "/")
	return bub
}

func (bub *BinanceURLBuilder) Ping() *BinanceURLBuilder {
	bub.clean()
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(PING),
	}, "/")
	return bub
}
