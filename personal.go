package taxcalc

import (
	"github.com/shopspring/decimal"
)

type Module interface {
	Load(map[string]string) error
	Calculate() (string, decimal.Decimal)
}

func PersonalTax() decimal.Decimal {
	//calculate taxable income

	//calculate credits
	return decimal.NewFromFloat(0.0)
}
