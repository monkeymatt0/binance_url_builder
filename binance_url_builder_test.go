package binance_url_builder

import (
	"testing"
)

const (
	testKlines           = "https://testnet.binance.vision/api/v3/klines"
	testKlinesWithParams = "https://testnet.binance.vision/api/v3/klines?interval=5m&limit=1000&symbol=BTCUSDT"
	testOrder            = "https://testnet.binance.vision/api/v3/order"
	testAccount          = "https://testnet.binance.vision/api/v3/account"
	testPing             = "https://testnet.binance.vision/api/v3/ping"

	prodKlines           = "https://api.binance.com/api/v3/klines"
	prodKlinesWithParams = "https://api.binance.com/api/v3/klines?interval=5m&limit=1000&symbol=BTCUSDT"
	prodOrder            = "https://api.binance.com/api/v3/order"
	prodAccount          = "https://api.binance.com/api/v3/account"
	prodPing             = "https://api.binance.com/api/v3/ping"
)

func TestNew(t *testing.T) {
	bub := &BinanceURLBuilder{}
	bub.New(true)
	if bub.Host != string(TEST) {
		t.Errorf("Error while creating new test binance builder\n")
	}
	bub.New(false)
	if bub.Host != string(PRODUCTION) {
		t.Errorf("Error while creating new production binance builder\n")
	}
}

func TestUrls(t *testing.T) {
	bub := &BinanceURLBuilder{}
	// Testnet
	bub.New(true)
	params := make(map[string]string)
	pass := bub.Klines(params).String() == testKlines
	if !pass {
		t.Errorf("[TEST] Error with klines endpoint building")
	}
	params["symbol"] = "BTCUSDT"
	params["interval"] = "5m"
	params["limit"] = "1000"
	pass = bub.Klines(params).String() == testKlinesWithParams
	if !pass {
		t.Errorf("[TEST] Error with klines endpoint building")
	}
	delete(params, "symbol")
	delete(params, "interval")
	delete(params, "limit")
	pass = bub.Order().String() == testOrder
	if !pass {
		t.Errorf("[TEST] Error with order endpoint building")
	}
	pass = bub.Ping().String() == testPing
	if !pass {
		t.Errorf("[TEST] Error with ping endpoint building")
	}
	pass = bub.Account().String() == testAccount
	if !pass {
		t.Errorf("[TEST] Error with account endpoint building")
	}
	// Production
	bub.New(false)
	pass = bub.Klines(params).String() == prodKlines
	if !pass {
		t.Errorf("[PROD] Error with klines endpoint building")
	}
	params["symbol"] = "BTCUSDT"
	params["interval"] = "5m"
	params["limit"] = "1000"
	pass = bub.Klines(params).String() == prodKlinesWithParams
	if !pass {
		t.Errorf("[TEST] Error with klines endpoint building")
	}
	delete(params, "symbol")
	delete(params, "interval")
	delete(params, "limit")
	pass = bub.Order().String() == prodOrder
	if !pass {
		t.Errorf("[PROD] Error with order endpoint building")
	}
	pass = bub.Ping().String() == prodPing
	if !pass {
		t.Errorf("[PROD] Error with ping endpoint building")
	}
	pass = bub.Account().String() == prodAccount
	if !pass {
		t.Errorf("[PROD] Error with account endpoint building")
	}

}
