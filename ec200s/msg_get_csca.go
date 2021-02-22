package ec200s

import "AT_GO/atSerial"

//获取短信中心号码
func (m ec200s) MsgGetCSCA() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CSCA?\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
