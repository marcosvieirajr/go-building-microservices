package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewRates(log hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: log, rates: map[string]float64{}}
	err := er.getRates()
	return er, err
}

func (e *ExchangeRates) getRates() error {
	url := "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	md := &cubes{}
	xml.NewDecoder(resp.Body).Decode(md)

	for _, d := range md.CubeData {
		r, err := strconv.ParseFloat(d.Rate, 64)
		if err != nil {
			return err
		}

		e.rates[d.Currency] = r
	}

	return nil
}

type cubes struct {
	CubeData []cube `xml:"Cube>Cube>Cube"`
}

type cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
