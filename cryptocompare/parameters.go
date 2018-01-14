package cryptocompare

import (
	"fmt"
	"strings"
)

type Parameter interface {
	getQueryName() string
	getValueString() string
}

type ParamFsym string

func (p ParamFsym) getQueryName() string {
	return "fsym"
}

func (p ParamFsym) getValueString() string {
	return fmt.Sprintf("%v", p)
}

type ParamTsyms []string

func (p ParamTsyms) getQueryName() string {
	return "tsyms"
}

func (p ParamTsyms) getValueString() string {
	return fmt.Sprintf("%v", combineMultiParam(p))
}

type ParamExchange string

func (p ParamExchange) getQueryName() string {
	return "e"
}

func (p ParamExchange) getValueString() string {
	return fmt.Sprintf("%v", p)
}

type ParamExtraparams string

func (p ParamExtraparams) getQueryName() string {
	return "extraParams"
}

func (p ParamExtraparams) getValueString() string {
	return fmt.Sprintf("%v", p)
}

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

func combineMultiParam(params []string) string {
	if params == nil || len(params) == 0 {
		return ""
	}

	return strings.Join(params, ",")
}
