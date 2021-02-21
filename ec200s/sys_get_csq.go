package ec200s

import "AT_GO/atSerial"

//查询信号强度，第一个值为0-31则正常，99为不正常
func (m ec200s) SysGetCSQ() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CSQ\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
