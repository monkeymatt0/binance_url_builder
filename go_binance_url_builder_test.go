package go_binance_url_builder

import "testing"

const (
	testKlines  = "https://testnet.binance.vision/api/v3/klines"
	testOrder   = "https://testnet.binance.vision/api/v3/order"
	testAccount = "https://testnet.binance.vision/api/v3/account"
	testPing    = "https://testnet.binance.vision/api/v3/ping"

	prodKlines  = "https://api.binance.com/api/v3/klines"
	prodOrder   = "https://api.binance.com/api/v3/order"
	prodAccount = "https://api.binance.com/api/v3/account"
	prodPing    = "https://api.binance.com/api/v3/ping"
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
	pass := bub.Klines().String() == testKlines
	if !pass {
		t.Errorf("[TEST] Error with klines endpoint building")
	}
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
	pass = bub.Klines().String() == prodKlines
	if !pass {
		t.Errorf("[PROD] Error with klines endpoint building")
	}
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
