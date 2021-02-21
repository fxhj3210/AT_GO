package ec200s

import (
	"AT_GO/atSerial"
)

//显示产品标识信息(显示全部）
func (m ec200s) SysGetATI() (CGSN string, err error) {

	results, err := atSerial.SendInstructions(m.Machine, []byte("ATI\r"))
	if err != nil {
		return "", err
	}

	return string(results), err
}
