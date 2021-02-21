package ec200s

import "AT_GO/atSerial"

//查询SIM卡是否被网络接受.
func (m ec200s) NetworkGetCOPS() (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("at+cops?\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
