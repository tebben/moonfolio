package cryptocompare

import (
	"fmt"
	"strings"
)

func buildURI(endpoint string, optionals OptionalParameters, parameters ...Parameter) string {
	uri := endpoint

	if optionals != nil {
		parameters = append(parameters, optionals.GetParameters()...)
	}

	for _, p := range parameters {
		if p == nil {
			continue
		}

		uri = appendParameter(uri, p)
	}

	return uri
}

func appendParameter(uri string, p Parameter) string {
	prefix := "?"
	if strings.Contains(uri, prefix) {
		prefix = "&"
	}

	return fmt.Sprintf("%s%s%s", uri, prefix, p.ToString())
}
