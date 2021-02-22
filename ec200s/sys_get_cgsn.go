package ec200s

import (
	"AT_GO/atSerial"
)

//请求产品序列号标识
func (m ec200s) SysGetCGSN() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CGSN\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
