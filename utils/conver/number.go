package conver

import "strconv"

func StrToUint32(str string) (uint32, error) {
	num, err := strconv.ParseUint(str, 10, 32)
	return uint32(num), err
}
