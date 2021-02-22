package ec200s

import "AT_GO/atSerial"

//删除短信
/*
AT+CMGD=<index>
	前面使用AT+CMGL可以列出短信列表，从而可以获取很多关于短信的信息，其中一项是index（短信编号），这个命令指删除index位置的短信
	例:AT+CMGD=0

AT+CMGD=<index>[,<delflag>]
	忽略index的值，根据的delflag的值来做处理
		<delflag>
		0 —— 删除<index>中指定的消息
		1 —— 删除<mem1>存储中的所有已读消息
		2 —— 删除<mem1>存储中的所有已读消息以及已发送的移动原始消息
		3 —— 删除<mem1>存储中的所有已读消息以及所有已发送和未发送的移动始发消息
		4 —— 删除<mem1>存储中的所有消息
	例:AT+CMGD=1,4
*/

func (m ec200s) MsgDelCMGD(Index string, DelFlag string) (results []byte, err error) {

	if DelFlag == "" {
		results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CMGD="+Index+"\r"))
	} else {
		results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CMGD="+Index+","+DelFlag+"\r"))
	}

	if err != nil {
		return nil, err
	}

	return results, err
}
