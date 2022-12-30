package storage

import (
	"github.com/onmono/rest-api-cryptocurrency/internal/types/cbr"
	"github.com/onmono/rest-api-cryptocurrency/internal/types/kucoin"
	"github.com/onmono/rest-api-cryptocurrency/internal/types/storage"
	"time"
)

type Storage interface {
	// TODO
	BTCUSDTStorage
	FiatStorage
}

type FiatStorage interface {
	GetFiatLast() (types.GetLastValueFiat, error)
	GetFiatHistory() (types.GetHistoryValuesFiat, error)
	GetFiatDate() (time.Time, error)
	InsertFiat(date time.Time, fiat cbr.Fiat) (bool, error)
}

type BTCUSDTStorage interface {
	GetBTCUSDTLast() (kucoin.BTCUSDT, error)
	GetBTCUSDTHistory() (types.HistoryValuesBTCUSDT, error)
	InsertBTCUSDT(btcUsdt kucoin.BTCUSDT) (bool, error)
}
