package cmc

var fiatSymbols []string

// Fiat type to use for fiat enumeration
type Fiat int

// String returns the string representation of the fiat enum
func (e Fiat) String() string {
	return fiatSymbols[int(e)]
}

// Symbols for available fiat currencies
var (
	FiatUSD = ciota("USD")
	FiatAUD = ciota("AUD")
	FiatBRL = ciota("BRL")
	FiatCAD = ciota("CAD")
	FiatCHF = ciota("CHF")
	FiatCLP = ciota("CLP")
	FiatCNY = ciota("CNY")
	FiatCZK = ciota("CZK")
	FiatDKK = ciota("DKK")
	FiatEUR = ciota("EUR")
	FiatGBP = ciota("GBP")
	FiatHKD = ciota("HKD")
	FiatHUF = ciota("HUF")
	FiatIDR = ciota("IDR")
	FiatILS = ciota("ILS")
	FiatINR = ciota("INR")
	FiatJPY = ciota("JPY")
	FiatKRW = ciota("KRW")
	FiatMXN = ciota("MXN")
	FiatMYR = ciota("MYR")
	FiatNOK = ciota("NOK")
	FiatNZD = ciota("NZD")
	FiatPHP = ciota("PHP")
	FiatPKR = ciota("PKR")
	FiatPLN = ciota("PLN")
	FiatRUB = ciota("RUB")
	FiatSEK = ciota("SEK")
	FiatSGD = ciota("SGD")
	FiatTHB = ciota("THB")
	FiatTRY = ciota("TRY")
	FiatTWD = ciota("TWD")
	FiatZAR = ciota("ZAR")
)

func ciota(s string) Fiat {
	fiatSymbols = append(fiatSymbols, s)
	return Fiat(len(fiatSymbols) - 1)
}

// GetSupportedFiat returns the string representations of
// the CMC supported fiat currencies used for converting
func GetSupportedFiat() []string {
	return fiatSymbols
}
