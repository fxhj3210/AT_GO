package ec200s

import "AT_GO/atSerial"

//获取当前短信格式,0是PDU模式，1是TEXT模式
func (m ec200s) MsgGetCMGF() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CMGF?\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
