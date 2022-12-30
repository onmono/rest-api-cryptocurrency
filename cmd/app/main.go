package main

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/onmono/rest-api-cryptocurrency/api"
	"github.com/onmono/rest-api-cryptocurrency/internal/service/btcusdt"
	"github.com/onmono/rest-api-cryptocurrency/internal/service/fiat"
	"github.com/onmono/rest-api-cryptocurrency/internal/storage"
	"log"
	"time"
)

const (
	fiatCbrUrl       = "http://www.cbr.ru/scripts/XML_daily.asp"
	btcUsdtKucoinUrl = "https://api.kucoin.com/api/v1/market/stats?symbol=BTC-USDT"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	fiatDuration := flag.Duration("duration", 5*time.Second, "timer duration for fetch fiat data")
	flag.Parse()

	store := storage.NewPostgreStorage()
	fiatService := fiat.NewService(fiatCbrUrl)
	btcUsdtService := btcusdt.NewService(btcUsdtKucoinUrl)

	server := api.NewServer(*listenAddr, store, fiatService, btcUsdtService)

	go server.FetchFiatJob(time.NewTimer(0), *fiatDuration)
	go server.FetchBTCUSDTJob(time.NewTimer(0), *fiatDuration)

	fmt.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())
}
