package ec200s

import "AT_GO/atSerial"

//请求产品序列号标识
func (m ec200s) SysGetCMEE() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CMEE?\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
