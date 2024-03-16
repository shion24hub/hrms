package bybit

import (
	"compress/gzip"
	"context"
	"encoding/csv"
	"errors"
	"net/http"
	"net/url"
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

func FetchTradingData(url string) (*csv.Reader, error) {

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

	return r, nil
}
