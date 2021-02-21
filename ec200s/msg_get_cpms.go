package ec200s

import "AT_GO/atSerial"

//获取短信存储载体,MT(终端),SM(SIM卡),ME(手机设备)
func (m ec200s) MsgGetCPMS() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CPMS=?\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
