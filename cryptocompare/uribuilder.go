package cryptocompare

import (
	"fmt"
	"strings"
)

func buildURI(endpoint string, parameters ...Parameter) string {
	uri := endpoint

	for _, p := range parameters {
		uri = appendParameter(uri, p)
	}

	// debug request uri
	//log.Printf("%s", uri)

	return uri
}

func appendParameter(uri string, p Parameter) string {
	if p == nil || p.getValueString() == "" {
		return uri
	}

	prefix := "?"
	if strings.Contains(uri, prefix) {
		prefix = "&"
	}

	return fmt.Sprintf("%s%s%s=%v", uri, prefix, p.getQueryName(), p.getValueString())
}
