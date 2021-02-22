package ec200s

import "AT_GO/mode"

/*
 发送短消息常用Text和PDU(Protocol   Data   Unit，协议数据单元)模式。使用Text模式收发短信代码简单，实现起来十分容易，但最大的缺点是不能收发中文短信；而PDU模式不仅支持中文短信，也能发送英文短信。PDU模式收发短信可以使用3种编码：7-bit、8-bit和UCS2编码。7-bit编码用于发送普通的ASCII字符，8-bit编码通常用于发送数据消息，UCS2编码用于发送Unicode字符。一般的PDU编码由A   B   C   D   E   F   G   H   I   J   K   L   M十三项组成。
	A：短信息中心地址长度，2位十六进制数(1字节)。
	B：短信息中心号码类型，2位十六进制数。
	C：短信息中心号码，B+C的长度将由A中的数据决定。
	D：文件头字节，2位十六进制数。
	E：信息类型，2位十六进制数。
	F：被叫号码长度，2位十六进制数。
	G：被叫号码类型，2位十六进制数，取值同B。
	H：被叫号码，长度由F中的数据决定。
	I：协议标识，2位十六进制数。
	J：数据编码方案，2位十六进制数。
	K：有效期，2位十六进制数。
	L：用户数据长度，2位十六进制数。
	M：用户数据，其长度由L中的数据决定。J中设定采用UCS2编码，这里是中英文的Unicode字符。
*/

func (m ec200s) MSGCreate(SMS mode.MSG) (binary []byte, err error) {

	return
}
