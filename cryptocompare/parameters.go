package cryptocompare

import (
	"fmt"
	"strings"
)

type OptionalParameters interface {
	GetParameters() []Parameter
}

type Parameter interface {
	ToString() string
}

type Fsym string

func (p *Fsym) ToString() string {
	return fmt.Sprintf("fsym=%s", *p)
}

type Fsyms []string

func (p *Fsyms) ToString() string {
	return fmt.Sprintf("fsyms=%s", combineMultiParam(*p))
}

type Tsym string

func (p *Tsym) ToString() string {
	return fmt.Sprintf("tsym=%s", *p)
}

type Tsyms []string

func (p *Tsyms) ToString() string {
	return fmt.Sprintf("tsyms=%s", combineMultiParam(*p))
}

type ExtraParams string

func (p *ExtraParams) ToString() string {
	return fmt.Sprintf("extraParams=%s", *p)
}

type Exchange string

func (p *Exchange) ToString() string {
	return fmt.Sprintf("e=%s", *p)
}

type Sign bool

func (p *Sign) ToString() string {
	return fmt.Sprintf("sign=%v", *p)
}

// TryConversion If set to false, it will try to get values without using any conversion at all
type TryConversion bool

func (p *TryConversion) ToString() string {
	return fmt.Sprintf("tryConversion=%v", *p)
}

func combineMultiParam(params []string) string {
	if params == nil || len(params) == 0 {
		return ""
	}

	return strings.Join(params, ",")
}

// GetPriceOptionals can be supplied to GetPrice to set some optional parameters
//  e: Name of exchange. Default: CCCAGG
//  extraParams:
//  sign: If set to true, the server will sign the requests.
//  tryConversion: If set to false, it will try to get values without using any conversion at all
type GetPriceOptionals struct {
	ExtraParams   *ExtraParams
	Exchange      *Exchange
	Sign          *Sign
	TryConversion *TryConversion
}

func (g *GetPriceOptionals) GetParameters() []Parameter {
	return []Parameter{g.Exchange, g.ExtraParams, g.Sign, g.TryConversion}
}
