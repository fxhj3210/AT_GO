package ec200s

import (
	"AT_GO/atSerial"
)

//读取单条短信
func (m ec200s) MsgGetCMGR(Index string) (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CMGR="+Index+"\r"))
	if err != nil {
		return nil, err
	}

	//扔去解码
	//serialization.MsgSerialization(results)

	return results, err
}
