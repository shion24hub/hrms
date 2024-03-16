package bybit

import (
	"compress/gzip"
	"context"
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type BybitTradingRow struct {
	Timestamp time.Time
	Symbol    string
	Side      string
	Size      string
	Price     string
}

func MakeUrl(symbol string, date time.Time) (string, error) {

	base := "https://public.bybit.com/trading/"
	fileName := symbol + date.Format("2006-01-02") + ".csv.gz"
	ep, err := url.JoinPath(base, symbol, fileName)
	if err != nil {
		return "", errors.New("failed to join path")
	}

	return ep, nil
}

func FetchTradingData(url string) ([]BybitTradingRow, error) {

	btrs := []BybitTradingRow{}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(gz)

	r.Read() // skip header
	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// convert unixms to time.Time
		unixms, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, err
		}
		ts := time.UnixMilli(int64(unixms * 1000))

		btr := BybitTradingRow{
			Timestamp: ts,
			Symbol:    row[1],
			Side:      row[2],
			Size:      row[3],
			Price:     row[4],
		}
		btrs = append(btrs, btr)
	}

	return btrs, nil
}
