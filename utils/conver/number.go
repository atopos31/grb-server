package conver

import "strconv"

func StrToUint32(str string) (uint32, error) {
	num, err := strconv.ParseUint(str, 10, 32)
	return uint32(num), err
}

func StrToUint8(str string) (uint8, error) {
	num, err := strconv.ParseUint(str, 10, 8)
	return uint8(num), err
}

func StrToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	return num, err
}

func StrToUInt(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 32)
	return uint(num), err
}
