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

func (bub *BinanceURLBuilder) Klines() *BinanceURLBuilder {
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(KLINES),
	}, "/")
	return bub
}

func (bub *BinanceURLBuilder) Order() *BinanceURLBuilder {
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(ORDER),
	}, "/")
	return bub
}

func (bub *BinanceURLBuilder) Account() *BinanceURLBuilder {
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(ACCOUNT),
	}, "/")
	return bub
}

func (bub *BinanceURLBuilder) Ping() *BinanceURLBuilder {
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(PING),
	}, "/")
	return bub
}
