package binance_url_builder

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type BinanceURLBuilder struct {
	url.URL
}

func (bub *BinanceURLBuilder) sign(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
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

func (bub *BinanceURLBuilder) UserDataStream(listenKey string) string {
	bub.clean()
	bub.Scheme = string(WSS)
	if bub.Host == string(TEST) {
		bub.Path = strings.Join([]string{
			string(WSS_TEST),
			string(WSS_BASE_PATH),
			listenKey,
		}, "/")
	} else {
		bub.Path = strings.Join([]string{
			string(WSS_PRODUCTION),
			string(WSS_BASE_PATH),
			listenKey,
		}, "/")
	}
	defer func() {
		bub.Scheme = string(HTTPS)
	}()
	return bub.String()
}

func (bub *BinanceURLBuilder) ListenKey(listenKey *string) *BinanceURLBuilder {
	bub.clean()
	if listenKey == nil {
		bub.Path = strings.Join([]string{
			string(BASE_PATH),
			string(LISTEN_KEY),
		}, "/")
	} else {
		bub.Path = strings.Join([]string{
			string(BASE_PATH),
			string(LISTEN_KEY),
		}, "/")
		query := bub.Query()
		query.Set("listenKey", *listenKey)
	}
	return bub
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

// @todo :
// Add params here to place an order properly
// Add param to asses if we need to delete the order
func (bub *BinanceURLBuilder) Order(params map[string]string, secret string) *BinanceURLBuilder {
	bub.clean()
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(ORDER),
	}, "/")
	query := bub.Query()
	// Create the query parameters for the POST
	for key, value := range params {
		query.Set(key, value)
	}
	encQuery := query.Encode()
	if secret != "" {
		signature := bub.sign(encQuery, secret)
		query.Set("signature", signature)
	}
	bub.RawQuery = query.Encode()
	return bub
}

func (bub *BinanceURLBuilder) Account(secret string) *BinanceURLBuilder {
	bub.clean()
	bub.Path = strings.Join([]string{
		string(BASE_PATH),
		string(ACCOUNT),
	}, "/")
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	query := bub.Query()
	if secret != "" {
		query.Set("timestamp", fmt.Sprintf("%d", timestamp))
		query.Set("omitZeroBalances", "true")
		encQuery := query.Encode()
		signature := bub.sign(encQuery, secret)
		query.Set("signature", signature)
	}
	bub.RawQuery = query.Encode()
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
