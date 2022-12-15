package api

import "net/http"

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/api/btcusdt", s.handleGetChangesBTCUSDT)
	http.HandleFunc("/api/currencies", s.handleGetFiatCurrenciesHistory)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetChangesBTCUSDT(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleGetFiatCurrenciesHistory(w http.ResponseWriter, r *http.Request) {

}
