package ec200s

import (
	"AT_GO/atSerial"
)

//请求产品序列号标识
func (m ec200s) SysGetCGSN() (CGSN string, err error) {

	results, err := atSerial.SendInstructions(m.Machine, []byte("AT+CGSN\r"))
	if err != nil {
		return "", err
	}

	return string(results), err
}
