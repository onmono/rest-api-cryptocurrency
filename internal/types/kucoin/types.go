package kucoin

import "time"

type BTCUSDT struct {
	Data `json:"data"`
}

type Data struct {
	Time time.Time `json:"time"`
	Buy  string    `json:"buy"`
	Sell string    `json:"sell"`
}

func NewBTCUSDT(time time.Time, buy string, sell string) BTCUSDT {
	return BTCUSDT{Data: Data{
		Time: time,
		Buy:  buy,
		Sell: sell,
	}}
}

func (b BTCUSDT) IsDifferent(that BTCUSDT) bool {
	return b.Data.Sell != that.Data.Sell || b.Data.Buy != that.Data.Buy
}
