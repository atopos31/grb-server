package errcode

import "fmt"

type Error struct {
	code int
	msg  string
}

var errorCodes = map[int]struct{}{}

func NewError(code int, msg string) Error {
	if _, ok := errorCodes[code]; ok {
		panic(fmt.Sprintf("code %d is exsit, please change one", code))
	}
	errorCodes[code] = struct{}{}
	return Error{code: code, msg: msg}
}

func (e Error) Error() string {
	return e.msg
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Msg() string {
	return e.msg
}
