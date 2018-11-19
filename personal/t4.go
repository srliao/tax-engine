package t4

import (
	"github.com/shopspring/decimal"
)

type T4 struct {
	income decimal.Decimal
}

func (t *T4) Load(data map[string]string) error {
	return nil
}

func (t *T4) Calculate() (string, decimal.Decimal) {
	return "", decimal.NewFromFloat(0.0)
}
