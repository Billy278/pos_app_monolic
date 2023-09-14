package helper

import "errors"

func CekReqType(data string) (err error) {
	if data == "CASH" || data == "E-WALLET" || data == "EDC" {
		return
	}
	err = errors.New("Bad type request")

	return
}
