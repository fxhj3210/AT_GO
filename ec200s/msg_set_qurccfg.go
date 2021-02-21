package ec200s

import "AT_GO/atSerial"

//设置短信URC输出的通道,分为："usbat"(USB的AT虚拟串口,默认),"usbmodem"(usb接入后，用于拨号上网的端口),"uart1"(硬件串口)
func (m ec200s) MsgSetQURCCFG(Port string) (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+QURCCFG="+"\"urcport\",\""+Port+"\""+"\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
