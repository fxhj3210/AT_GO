package ec200s

import "AT_GO/atSerial"

//设置短信存储载体,MT(终端),SM(SIM卡),ME(手机设备)
func (m ec200s) MsgSetCPMS(a string, b string, c string) (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CPMS="+"\""+a+"\",\""+b+"\",\""+c+"\""+"\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
