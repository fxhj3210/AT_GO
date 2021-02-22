package ec200s

import "AT_GO/atSerial"

//设置短信格式,0是PDU模式，1是TEXT模式
func (m ec200s) MsgSetCMGF(C string) (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CMGF="+C+"\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
