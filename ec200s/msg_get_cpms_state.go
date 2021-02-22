package ec200s

import "AT_GO/atSerial"

//获取当前短信存储设置
func (m ec200s) MsgGetCPMSState() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CPMS?\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
