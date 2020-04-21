package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func AccRandTimeInt() int {
	timeStr := strings.Replace(time.Now().Local().Format("060102150405.00"), ".", "", 1)
	postfix := acctest.RandStringFromCharSet(4, "0123456789")

	i, err := strconv.Atoi(timeStr + postfix)
	if err != nil {
		panic(err)
	}

	return i
}
