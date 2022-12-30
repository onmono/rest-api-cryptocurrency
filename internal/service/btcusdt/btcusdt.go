package btcusdt

import (
	"context"
	"encoding/json"
	"github.com/onmono/rest-api-cryptocurrency/internal/types/kucoin"
	"net/http"
	"time"
)

type Service interface {
	GetBtcUsdtData(context.Context) (kucoin.BTCUSDT, error)
}

type service struct {
	url string
}

func NewService(url string) Service {
	return &service{
		url: url,
	}
}

type BTCUSDT struct {
	Data `json:"data"`
}

type Data struct {
	Time int64  `json:"time"` // Unix Millisecond
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
}

func (s *service) GetBtcUsdtData(ctx context.Context) (data kucoin.BTCUSDT, err error) {
	temp := BTCUSDT{}
	resp, err := http.Get(s.url)
	if err != nil {
		return kucoin.BTCUSDT{}, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&temp); err != nil {
		return kucoin.BTCUSDT{}, err
	}
	data = kucoin.NewBTCUSDT(time.UnixMilli(temp.Data.Time), temp.Buy, temp.Sell)

	//price, err := strconv.ParseFloat(data.Data.AveragePrice, 10)
	//fmt.Println("price:", price)
	return data, nil
}
