package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type GetLastValueBTCUSDT struct {
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Value     float64   `bson:"value" json:"value"`
}

type HistoryValuesBTCUSDT struct {
	Total          int                   `json:"total"`
	LastValBTCUSDT []GetLastValueBTCUSDT `json:"history"`
}

type LastValueFiat struct {
	// TODO: format Date like YYYY-MM-DD
	Date time.Time `bson:"date" json:"date"`
	Data `bson:"data" json:"data"`
}

type Data map[string]interface{}

func (a Data) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Data) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type GetHistoryValuesFiat struct {
	Total            int                `json:"total"`
	GetLastValueFiat []GetLastValueFiat `json:"history"`
}

type GetLastValueFiat struct {
	// TODO: format Date like YYYY-MM-DD
	Date string `json:"date"`
	Data `json:"data"`
}
