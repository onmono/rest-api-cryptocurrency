package fiat

import (
	"context"
	"encoding/xml"
	"github.com/onmono/rest-api-cryptocurrency/internal/types/cbr"
	"golang.org/x/net/html/charset"
	"net/http"
	"strings"
)

type Service interface {
	GetData(context.Context) (cbr.Fiat, error)
}

type service struct {
	url string
}

func NewService(url string) Service {
	return &service{
		url: url,
	}
}

func (s *service) GetData(ctx context.Context) (data cbr.Fiat, err error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return cbr.Fiat{}, err
	}

	defer resp.Body.Close()
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&data); err != nil {
		return cbr.Fiat{}, err
	}
	for _, v := range data.Valute {
		v.Value = strings.Replace(v.Value, ",", ".", -1)
		_ = v.Value
	}
	return data, nil
}
