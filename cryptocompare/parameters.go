package cryptocompare

import (
	"fmt"
	"strings"
)

// Parameter that can be passed to CryptoCompare
type Parameter interface {
	getQueryName() string
	getValueString() string
}

// ParamFsym From Symbol
type ParamFsym string

func (p ParamFsym) getQueryName() string {
	return "fsym"
}

func (p ParamFsym) getValueString() string {
	return fmt.Sprintf("%v", p)
}

// ParamFsyms from Symbols, include multiple symbols
type ParamFsyms []string

func (p ParamFsyms) getQueryName() string {
	return "fsyms"
}

func (p ParamFsyms) getValueString() string {
	return fmt.Sprintf("%v", combineMultiParam(p))
}

// ParamTsym to Symbol
type ParamTsym string

func (p ParamTsym) getQueryName() string {
	return "tsym"
}

func (p ParamTsym) getValueString() string {
	return fmt.Sprintf("%v", p)
}

// ParamTsyms to Symbols, include multiple symbols
type ParamTsyms []string

func (p ParamTsyms) getQueryName() string {
	return "tsyms"
}

func (p ParamTsyms) getValueString() string {
	return fmt.Sprintf("%v", combineMultiParam(p))
}

// ParamExchange name of exchange. Default: CCCAGG
type ParamExchange string

func (p ParamExchange) getQueryName() string {
	return "e"
}

func (p ParamExchange) getValueString() string {
	return fmt.Sprintf("%v", p)
}

// ParamExtraparams Name of your application, no default
type ParamExtraparams string

func (p ParamExtraparams) getQueryName() string {
	return "extraParams"
}

func (p ParamExtraparams) getValueString() string {
	return fmt.Sprintf("%v", p)
}

// ParamSign If set to true, the server will sign the requests, default false
type ParamSign bool

func (p ParamSign) getQueryName() string {
	return "sign"
}

func (p ParamSign) getValueString() string {
	if p == false {
		return ""
	}

	return fmt.Sprintf("%v", p)
}

// ParamTryConversion if set to false, it will try to get values without using any conversion at all
type ParamTryConversion bool

func (p ParamTryConversion) getQueryName() string {
	return "tryConversion"
}

func (p ParamTryConversion) getValueString() string {
	if p == false {
		return ""
	}

	return fmt.Sprintf("%v", p)
}

// ParamAggregate used to get aggregated data from histo endpoints
type ParamAggregate int

func (p ParamAggregate) getQueryName() string {
	return "aggregate"
}

func (p ParamAggregate) getValueString() string {
	if p == 0 {
		return ""
	}

	return fmt.Sprintf("%v", p)
}

// ParamLimit limit response by limit amount, max 2000
type ParamLimit int

func (p ParamLimit) getQueryName() string {
	return "limit"
}

func (p ParamLimit) getValueString() string {
	if p == 0 {
		return ""
	}

	return fmt.Sprintf("%v", p)
}

func combineMultiParam(params []string) string {
	if params == nil || len(params) == 0 {
		return ""
	}

	return strings.Join(params, ",")
}
