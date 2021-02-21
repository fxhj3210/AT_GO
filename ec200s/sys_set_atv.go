package ec200s

import (
	"AT_GO/atSerial"
)

//设置TA响应模式,0时：0,1时：ok
func (m ec200s) SysSetATV(Model string) (CGSN string, err error) {

	results, err := atSerial.SendInstructions(m.Machine, []byte("ATV"+Model+"\r"))
	if err != nil {
		return "", err
	}

	return string(results), err
}
