package ec200s

import "AT_GO/atSerial"

//读取短信,("REC UNREAD","REC READ","STO UNSENT","STO SENT","ALL")
//(“REC UNREAD”-未读;“REC READ”-已读;“STO UNSENT”-待发;“STO SENT”-已发;“ALL”-全部)
func (m ec200s) MsgGetCMGL(C string) (results []byte, err error) {

	results, err = atSerial.SendInstructions(m.Machine, []byte("AT+CMGL=\""+C+"\"\r"))
	if err != nil {
		return nil, err
	}

	return results, err
}
