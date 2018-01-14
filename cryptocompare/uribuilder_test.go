package cryptocompare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestParam string

func (t *TestParam) ToString() string {
	return fmt.Sprintf("test=%s", *t)
}

type TestOptionals struct {
	ExtraParams   *ExtraParams
	Exchange      *Exchange
	Sign          *Sign
	TryConversion *TryConversion
}

func (g *TestOptionals) GetParameters() []Parameter {
	return []Parameter{g.Exchange, g.ExtraParams, g.Sign, g.TryConversion}
}

func TestBuildUri(t *testing.T) {
	endpoint := "https://coins?user=barry"
	testParam := TestParam("test")
	a := &testParam
	uri := buildURI(endpoint, nil, a)

	assert.Equal(t, "https://coins?user=barry&test=test", uri)
}

func TestBuildUriWithOptionals(t *testing.T) {
	endpoint := "https://coins"
	testParam := TestParam("test")

	exchange := Exchange("Kraken")
	sign := Sign(true)
	optionals := TestOptionals{Exchange: &exchange, Sign: &sign}

	uri := buildURI(endpoint, &optionals, *testParam)

	assert.Equal(t, "https://coins?test=test&e=Kraken&sign=true", uri)
}

func TestAppendParameter(t *testing.T) {
	uri := "https://coins"
	testParam1 := TestParam("test")
	testParam2 := TestParam("ikel")

	uri = appendParameter(uri, &testParam1)
	uri = appendParameter(uri, &testParam2)

	assert.Equal(t, "https://coins?test=test&test=ikel", uri)
}
