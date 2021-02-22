package ec200s

import "AT_GO/atSerial"

//获取国际移动用户识别码
func (m ec200s) SysGetCIMI() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CIMI\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
