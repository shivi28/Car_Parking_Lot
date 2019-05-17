package lib

import (
	"strconv"
	"strings"
)

// this converts the string to integer
func ConvertStringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
