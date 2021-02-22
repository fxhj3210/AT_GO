package ec200s

import "AT_GO/atSerial"

//设置错误讯息格式,0禁用结果代码;1启用结果代码并使用数值;2启用结果代码并使用详细值
func (m ec200s) SysSetCMEE(Model string) (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CMEE="+Model+"\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
