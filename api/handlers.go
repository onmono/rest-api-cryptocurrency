package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleChangesBTCUSDT(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetChangesBTCUSDT(w, r)
	case http.MethodPost:
		s.handlePostChangesBTCUSDT(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetChangesBTCUSDT(w http.ResponseWriter, r *http.Request) {
	// TODO: вывести последнее текущее значение

	data, err := s.store.GetBTCUSDTLast()
	if err != nil {
		err := writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		if err != nil {
			return
		}
		return
	}

	err = writeJSON(w, http.StatusOK, data)
	if err != nil {
		return
	}
}

func (s *Server) handlePostChangesBTCUSDT(w http.ResponseWriter, r *http.Request) {
	// TODO: вывести историю с фильтрами по дате и времени и пагинацией

	data, err := s.store.GetBTCUSDTHistory()
	if err != nil {
		err := writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		if err != nil {
			return
		}
		return
	}

	err = writeJSON(w, http.StatusOK, data)
	if err != nil {
		return
	}
}

func (s *Server) handleFiatCurrenciesHistory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetFiatCurrenciesHistory(w, r)
	case http.MethodPost:
		s.handlePostFiatCurrenciesHistory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetFiatCurrenciesHistory(w http.ResponseWriter, r *http.Request) {
	// TODO: отобразить последние (текущие) курсы фиатных валют по отношению к рублю
	data, err := s.store.GetFiatLast()
	if err != nil {
		err := writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		if err != nil {
			return
		}
		return
	}

	err = writeJSON(w, http.StatusOK, data)
	if err != nil {
		return
	}
}

func (s *Server) handlePostFiatCurrenciesHistory(w http.ResponseWriter, r *http.Request) {
	// TODO: возвратить историю изменения с фильтрами по дате и валюте и пагинацией
	data, err := s.store.GetFiatHistory()
	if err != nil {
		err := writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		if err != nil {
			return
		}
		return
	}

	err = writeJSON(w, http.StatusOK, data)
	if err != nil {
		return
	}
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
