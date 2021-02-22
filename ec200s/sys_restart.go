package ec200s

import "AT_GO/atSerial"

//重启模块,会断电
func (m ec200s) SysRestart() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CFUN=1,1\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
