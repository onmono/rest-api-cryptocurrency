package api

import (
	"context"
	"fmt"
	"github.com/onmono/rest-api-cryptocurrency/internal/service/btcusdt"
	"github.com/onmono/rest-api-cryptocurrency/internal/service/fiat"
	"github.com/onmono/rest-api-cryptocurrency/internal/storage"
	"net/http"
	"time"
)

type Server struct {
	listenAddr     string
	store          storage.Storage
	fiatService    fiat.Service
	btcUsdtService btcusdt.Service
}

func NewServer(listenAddr string, store storage.Storage, fiatService fiat.Service, btcUsdtService btcusdt.Service) *Server {
	return &Server{
		listenAddr:     listenAddr,
		store:          store,
		fiatService:    fiatService,
		btcUsdtService: btcUsdtService,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/api/btcusdt", s.handleChangesBTCUSDT)
	http.HandleFunc("/api/currencies", s.handleFiatCurrenciesHistory)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) FetchFiatJob(t *time.Timer, duration time.Duration) {
	for {
		<-t.C
		var err error
		lastDate, err := s.store.GetFiatDate()
		if err != nil {
			// TODO
		}

		now := time.Now()
		if isToday(lastDate) {
			t.Reset(duration)
			continue
		}
		data, err := s.fiatService.GetData(context.Background())
		if err != nil {
			//TODO
		}
		fmt.Printf("%v\n", data)
		_, err = s.store.InsertFiat(now, data)
		if err != nil {
			// TODO
		}
		t.Reset(duration)
	}
}

func isToday(that time.Time) bool {
	now := time.Now()
	return now.Format("2006-01-02") == that.Format("2006-01-02")
}

func (s *Server) FetchBTCUSDTJob(t *time.Timer, duration time.Duration) {
	for {
		<-t.C
		data, err := s.btcUsdtService.GetBtcUsdtData(context.Background())
		if err != nil {
			return
		}
		// проверить с курсом в бд, если изменился, то добавить новую строчку в бд
		valueBTCUSDT, err := s.store.GetBTCUSDTLast()
		var noRow bool
		if err != nil {
			if err.Error() == "GetBTCUSDTLast: no rows" {
				noRow = true
			} else {
				return
			}
		}

		if valueBTCUSDT.IsDifferent(data) || noRow {
			fmt.Println("valueBTCUSDT != data")
			_, err = s.store.InsertBTCUSDT(data)
			if err != nil {
				return
			}
		}
		fmt.Println("valueBTCUSDT == data")
		t.Reset(duration)
		//s.store.Query("INSERT INTO public.btc_usdt (timestamp,val) VALUES (now(),$1);", price)

		//return data, nil
	}
}
