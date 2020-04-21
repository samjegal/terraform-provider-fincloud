package features

import (
	"os"
	"strings"
)

func SupportsCustomTimeouts() bool {
	return strings.EqualFold(os.Getenv("FINCLOUD_CUSTOM_TIMEOUTS"), "true")
}
