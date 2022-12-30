package storage

import (
	"database/sql"
	"fmt"
	"github.com/onmono/rest-api-cryptocurrency/internal/types/cbr"
	"github.com/onmono/rest-api-cryptocurrency/internal/types/kucoin"
	"github.com/onmono/rest-api-cryptocurrency/internal/types/storage"
	"log"
	"reflect"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"

	driverName = "postgres"
)

type PostgreStorage struct {
	db *sql.DB
}

func NewPostgreStorage() *PostgreStorage {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open(driverName, psqlInfo)
	if err != nil {
		panic(err)
	}
	return &PostgreStorage{db: db}
}

func (r *PostgreStorage) GetLastBTCUSDT() (types.GetLastValueBTCUSDT, error) {
	data := types.GetLastValueBTCUSDT{}
	if err := r.db.QueryRow("SELECT value, timestamp FROM btc_usdt ORDER BY timestamp DESC LIMIT 1;").Scan(&data.Value, &data.Timestamp); err != nil {
		if err == sql.ErrNoRows {
			return types.GetLastValueBTCUSDT{}, fmt.Errorf("GetLastTimestampValBTCUSDT: no rows")
		}
		return types.GetLastValueBTCUSDT{}, fmt.Errorf("GetLastTimestampValBTCUSDT: err")
	}
	return data, nil
}

func (r *PostgreStorage) GetBTCUSDTHistory() (types.HistoryValuesBTCUSDT, error) {
	btcUsdtVals := make([]types.GetLastValueBTCUSDT, 0, 1)

	rows, err := r.db.Query("SELECT timestamp, value  FROM btc_usdt ORDER BY timestamp DESC;")
	if err != nil {
		if err == sql.ErrNoRows {
			return types.HistoryValuesBTCUSDT{}, fmt.Errorf("GetLastTimestampValBTCUSDT: no rows")
		}
		return types.HistoryValuesBTCUSDT{}, fmt.Errorf("GetLastTimestampValBTCUSDT: err")
	}
	counter := 0
	for rows.Next() {
		val := types.GetLastValueBTCUSDT{}

		s := reflect.ValueOf(&val).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)

		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}
		counter++
		log.Println(val)

		btcUsdtVals = append(btcUsdtVals, val)
	}

	data := types.HistoryValuesBTCUSDT{
		Total:          counter,
		LastValBTCUSDT: btcUsdtVals,
	}

	return data, nil
}

func (r *PostgreStorage) GetFiatLast() (types.GetLastValueFiat, error) {
	var data types.Data
	var date time.Time
	if err := r.db.QueryRow("SELECT date, data FROM fiats ORDER BY date DESC LIMIT 1;").
		Scan(&date, &data); err != nil {
		if err == sql.ErrNoRows {
			return types.GetLastValueFiat{}, fmt.Errorf("GetLastTimestampValBTCUSDT: no rows")
		}
		return types.GetLastValueFiat{}, fmt.Errorf("GetLastTimestampValBTCUSDT: err")
	}
	val := types.GetLastValueFiat{
		Date: date.Format("2006-01-02"),
		Data: data,
	}
	return val, nil
}

func (r *PostgreStorage) GetFiatHistory() (types.GetHistoryValuesFiat, error) {
	btcUsdtVals := make([]types.GetLastValueFiat, 0, 1)

	rows, err := r.db.Query("SELECT date, data  FROM fiats ORDER BY date DESC;")
	if err != nil {
		if err == sql.ErrNoRows {
			return types.GetHistoryValuesFiat{}, fmt.Errorf("GetLastTimestampValBTCUSDT: no rows")
		}
		return types.GetHistoryValuesFiat{}, fmt.Errorf("GetLastTimestampValBTCUSDT: err")
	}
	counter := 0

	for rows.Next() {
		var data types.Data
		var date time.Time
		err := rows.Scan(&date, &data)
		if err != nil {
			log.Fatal(err)
		}
		counter++
		val := types.GetLastValueFiat{
			Date: date.Format("2006-04-02"),
			Data: data,
		}

		log.Println(val)

		btcUsdtVals = append(btcUsdtVals, val)
	}

	data := types.GetHistoryValuesFiat{
		Total:            counter,
		GetLastValueFiat: btcUsdtVals,
	}

	return data, nil
}

func (r *PostgreStorage) GetFiatDate() (time.Time, error) {
	var date time.Time
	if err := r.db.QueryRow("SELECT date FROM fiats ORDER BY date DESC LIMIT 1;").
		Scan(&date); err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, fmt.Errorf("GetLastTimestampValBTCUSDT: no rows")
		}
		return time.Time{}, fmt.Errorf("GetLastTimestampValBTCUSDT: err")
	}
	return date, nil
}

func (r *PostgreStorage) InsertFiat(date time.Time, fiat cbr.Fiat) (bool, error) {
	bytes, _ := fiat.Marshal()
	sqlStatement := `INSERT INTO fiats (date, data) VALUES ($1, $2)`
	_, err := r.db.Exec(sqlStatement, date, bytes)
	if err != nil {
		panic(err)
	}
	return true, nil
}

func (r *PostgreStorage) InsertBTCUSDT(btcUsdt kucoin.BTCUSDT) (bool, error) {
	sqlStatement := `INSERT INTO public.btc_usdt (timestamp,buy,sell) VALUES ($1,$2,$3);`
	// TODO: Fix Bug
	_, err := r.db.Exec(sqlStatement, btcUsdt.Time, btcUsdt.Buy, btcUsdt.Sell)
	if err != nil {
		panic(err)
	}
	return true, nil
}

func (r *PostgreStorage) GetBTCUSDTLast() (kucoin.BTCUSDT, error) {
	var data kucoin.BTCUSDT
	if err := r.db.QueryRow("SELECT buy,sell FROM btc_usdt ORDER BY timestamp DESC LIMIT 1;").
		Scan(&data.Data.Buy, &data.Data.Sell); err != nil {
		if err == sql.ErrNoRows {
			return kucoin.BTCUSDT{}, fmt.Errorf("GetBTCUSDTLast: no rows")
		}
		return kucoin.BTCUSDT{}, fmt.Errorf("GetBTCUSDTLast: err")
	}
	return data, nil
}
