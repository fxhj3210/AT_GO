package ec200s

import "AT_GO/atSerial"

//查询当前SIM卡手机号
func (m ec200s) SysGetCNUM() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CNUM\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
