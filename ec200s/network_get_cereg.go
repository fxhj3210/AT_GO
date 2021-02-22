package ec200s

import "AT_GO/atSerial"

//查询模组是否注册上 NB 网络，0 未注册，MT 当前当前没有搜索或者注册到运营商网络;1 已注册，注册到归属网络;2 未注册;3 注册被拒绝;4 未知错误;5 已注册，注册到漫游网络
func (m ec200s) NetworkGetCEREG() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CEREG?\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
