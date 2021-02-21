package ec200s

import "AT_GO/atSerial"

//查询模组是否注册上GPRS网络，+CGREG:0,1 表示已注册上本地网，+CGREG:0,5表示注册上漫游网。
func (m ec200s) NetworkGetCGREG() (CGSN string, err error) {

	results, err := atSerial.SendInstructions(m.Machine, []byte("AT+CGREG?\r"))
	if err != nil {
		return "", err
	}

	return string(results), err
}
