package utils

import "strconv"

func ParseToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ParseToInt64(s string) (int64, error) {
	num, err := ParseToInt(s)
	return int64(num), err
}
