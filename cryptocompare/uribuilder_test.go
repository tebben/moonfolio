package cryptocompare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildUriWithFilledAndDefaultAndEmptyValues(t *testing.T) {
	//prepare
	endpoint := "https://coins"
	exchange := ParamExchange("Kraken")
	extraParams := ParamExtraparams("")
	sign := ParamSign(true)
	conversion := ParamTryConversion(false)

	// act
	uri := buildURI(endpoint, exchange, extraParams, sign, conversion)

	// assert
	assert.Equal(t, "https://coins?e=Kraken&sign=true", uri)
}

func TestAppendParameter(t *testing.T) {
	// prepare
	endpoint := "https://coins"
	toAppend := ParamExchange("Kraken")

	// act
	uri := appendParameter(endpoint, toAppend)

	// assert
	assert.Equal(t, fmt.Sprintf("https://coins?%s=%s", toAppend.getQueryName(), toAppend.getValueString()), uri)
}

func TestAppendParameterContainingQuestionMark(t *testing.T) {
	// prepare
	endpoint := "https://coins?s=1"
	toAppend := ParamExchange("Kraken")

	// act
	uri := appendParameter(endpoint, toAppend)

	// assert
	assert.Equal(t, fmt.Sprintf("https://coins?s=1&%s=%s", toAppend.getQueryName(), toAppend.getValueString()), uri)
}
