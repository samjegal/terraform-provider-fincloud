package common

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/Azure/go-autorest/autorest"
)

func GetQuery(queryParameters map[string]interface{}) string {
	parameters := autorest.MapToValues(queryParameters)
	v := make(url.Values, 0)
	for key, value := range parameters {
		for i := range value {
			d, err := url.QueryUnescape(value[i])
			if err != nil {
				return ""
			}
			value[i] = d
		}
		v[key] = value
	}
	return v.Encode()
}

func GetPath(baseUrl string, url string) string {
	r := regexp.MustCompile("https://[^{}/ :\\\\]+(?::\\d+)?")
	return r.ReplaceAllString(baseUrl, "") + url
}

func ensureValueString(value interface{}) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func ensureValueStrings(mapOfInterface map[string]interface{}) map[string]string {
	mapOfStrings := make(map[string]string)
	for key, value := range mapOfInterface {
		mapOfStrings[key] = ensureValueString(value)
	}
	return mapOfStrings
}

func GetPathParameters(baseUrl string, path string, pathParameters map[string]interface{}) string {
	parameters := ensureValueStrings(pathParameters)
	for key, value := range parameters {
		path = strings.Replace(path, "{"+key+"}", value, -1)
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return baseUrl + path
}
