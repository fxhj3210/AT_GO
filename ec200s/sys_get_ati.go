package ec200s

import (
	"AT_GO/atSerial"
)

//显示产品标识信息(显示全部）
func (m ec200s) SysGetATI() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("ATI\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
