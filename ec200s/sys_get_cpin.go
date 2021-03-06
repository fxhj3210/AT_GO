package ec200s

import (
	"AT_GO/atSerial"
)

//查询 SIM 卡状态，返回 READY 则表示SIM卡正常
func (m ec200s) SysGetCPIN() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CPIN?\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
