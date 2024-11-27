package binance_url_builder

// urls.go will contain the needed urls
// These urls are the ones needed to be used in the core trading bot

type Scheme string

const (
	HTTPS Scheme = "https"
	WSS   Scheme = "wss"
)

type BinanceHost string

const (
	PRODUCTION     BinanceHost = "api.binance.com"
	TEST           BinanceHost = "testnet.binance.vision"
	WSS_PRODUCTION BinanceHost = "stream.binance.com:9443"
	WSS_TEST       BinanceHost = "stream.testnet.binance.vision:9443"
)

type BinanceBasePath string

const (
	BASE_PATH     BinanceBasePath = "api/v3"
	WSS_BASE_PATH BinanceBasePath = "ws"
)

type BinancePath string

const (
	KLINES     BinancePath = "klines"
	ORDER      BinancePath = "order" // This can be used to CREATE, DELETE and CHECK the order
	PING       BinancePath = "ping"
	ACCOUNT    BinancePath = "account"
	LISTEN_KEY BinancePath = "userDataStream"
)
