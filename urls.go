package go_binance_url_builder

// urls.go will contain the needed urls
// These urls are the ones needed to be used in the core trading bot

type Scheme string

const (
	HTTPS Scheme = "https"
)

type BinanceHost string

const (
	PRODUCTION BinanceHost = "api.binance.com"
	TEST       BinanceHost = "testnet.binance.vision"
)

type BinanceBasePath string

const (
	BASE_PATH BinanceBasePath = "api/v3"
)

type BinancePath string

const (
	KLINES  BinancePath = "klines"
	ORDER   BinancePath = "order" // This can be used to CREATE, DELETE and CHECK the order
	PING    BinancePath = "ping"
	ACCOUNT BinancePath = "account"
)
