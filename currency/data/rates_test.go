package data

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestNewRates(t *testing.T) {
	r, err := NewRates(hclog.Default())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Rates: %#v\n", r.rates)
}
