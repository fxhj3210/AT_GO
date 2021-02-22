package atSerial

import (
	"AT_GO/mode"
	"errors"
	"fmt"
	"github.com/tarm/serial"
	"strings"
)

const MAXRWLEN = 8000
const SUCCATQ = "\r\nOK\r\n"
const FAILCATQ = "ERROR"

func SendInstructions(M mode.Machine, Command []byte) (results []byte, err error) {

	fmt.Println(string(Command))
	iorwc, err := serial.OpenPort(&serial.Config{
		Name:        M.ComName,
		Baud:        M.Baud,
		ReadTimeout: M.ReadTimeout,
		Size:        M.Size,
		Parity:      M.Parity,
		StopBits:    M.StopBits,
	})
	if err != nil {
		return nil, err
	}

	defer iorwc.Close()
	buffer := make([]byte, MAXRWLEN)

	//发命令之前清空缓冲区
	num, err := iorwc.Read(buffer)
	if err != nil {
		return nil, err
	}

	//发命令数据类型为[]byte
	num, err = iorwc.Write(Command)
	if err != nil {
		return nil, err
	}

	var tmpstr string = ""
	for i := 0; i < 3000; i++ {

		num, err = iorwc.Read(buffer)
		if num > 0 {
			tmpstr += fmt.Sprintf("%s", string(buffer[:num]))
		}

		//成功跳出处理
		if strings.LastIndex(tmpstr, SUCCATQ) > 0 {
			break
		}

		//失败返回错误
		if strings.LastIndex(tmpstr, FAILCATQ) > 0 {
			return nil, errors.New(FAILCATQ)
		}
	}

	//输出信息
	before := len(Command) + 2
	after := len(tmpstr) - len(SUCCATQ) - 2

	if before >= after {
		if after <= 0 {
			return nil, errors.New(FAILCATQ + ": No Data")
		}
		return []byte(tmpstr)[before : after+len(SUCCATQ)], err
	}
	return []byte(tmpstr)[before:after], err
}
