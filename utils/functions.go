package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(day int) string {
	dat, err := os.ReadFile(fmt.Sprintf("input/day%d.txt", day))
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(dat), "\r\n")
}
