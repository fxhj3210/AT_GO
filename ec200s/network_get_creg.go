package ec200s

import "AT_GO/atSerial"

//查询模组是否注册上GSM网络.如果返回 1 或 5 ，代表 CS 服务注册成功。1 表示已注册上本地网,5表示注册上漫游网
func (m ec200s) NetworkGetCREG() (CGSN string, err error) {

	results, err := atSerial.SendInstructions(m.Machine, []byte("AT+CREG?\r"))
	if err != nil {
		return "", err
	}
	return string(results), err
}
