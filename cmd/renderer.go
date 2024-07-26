package cmd

import (
	"strconv"
	"strings"
)

func createFormat(width, height int) string {
	var format strings.Builder

	format.WriteString("P3\n" + strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n255\n")
	return format.String()
}
