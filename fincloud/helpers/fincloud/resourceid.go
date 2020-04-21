package fincloud

import (
	"fmt"
	"net/url"
	"strings"
)

type ResourceID struct {
	Provider string
	Path     map[string]string
}

func ParseResourceId(id string) (*ResourceID, error) {
	url, err := url.ParseRequestURI(id)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse Financial Cloud ID: %s", err)
	}

	path := url.Path

	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")

	components := strings.Split(path, "/")

	if len(components)%2 != 0 {
		return nil, fmt.Errorf("The number of path segments is not divisible by 2 in %q", path)
	}

	componentMap := make(map[string]string, len(components)/2)
	for current := 0; current < len(components); current += 2 {
		key := components[current]
		value := components[current+1]

		if key == "" || value == "" {
			return nil, fmt.Errorf("Key/Value cannot be empty strings. Key: '%s', Value: '%s'", key, value)
		}

		componentMap[key] = value
	}

	obj := &ResourceID{}
	obj.Path = componentMap

	return obj, nil
}
