package ec200s

import (
	"fmt"
	"testing"
)

const Sub = "\n----------------------------"

func Test(t *testing.T) {
	machine := New("COM15")

	//设置TA响应模式,0时：0,1时：ok
	ATV, err := machine.SysSetATV("1")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(ATV), Sub)
	}

	//显示产品标识信息(显示全部）
	ATI, err := machine.SysGetATI()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(ATI), Sub)
	}

	//请求产品序列号标识
	CGSN, err := machine.SysGetCGSN()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CGSN), Sub)
	}

	//查询 SIM 卡状态，返回 READY 则表示SIM卡正常
	CPIN, err := machine.SysGetCPIN()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CPIN), Sub)
	}

	//查询信号强度，第一个值为0-31则正常，99为不正常
	CSQ, err := machine.SysGetCSQ()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CSQ), Sub)
	}

	//查询模组是否注册上GSM网络.如果返回 1 或 5 ，代表 CS 服务注册成功。1 表示已注册上本地网,5表示注册上漫游网
	CREG, err := machine.NetworkGetCREG()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CREG), Sub)
	}

	//查询模组是否注册上GPRS网络，+CGREG:0,1 表示已注册上本地网，+CGREG:0,5表示注册上漫游网。
	CGREG, err := machine.NetworkGetCGREG()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CGREG), Sub)
	}

	//查询SIM卡是否被网络接受.
	COPS, err := machine.NetworkGetCOPS()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(COPS), Sub)
	}

	//获取短信存储载体,MT(终端),SM(SIM卡),ME(手机设备)
	CPMS, err := machine.MsgGetCPMS()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CPMS), Sub)
	}

	//获取当前短信存储设置
	CPMSState, err := machine.MsgGetCPMSState()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CPMSState), Sub)
	}

	//设置短信存储载体,MT(终端),SM(SIM卡),ME(手机设备)
	MsgSetCPMS, err := machine.MsgSetCPMS("SM", "SM", "SM")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(MsgSetCPMS), Sub)
	}

	//设置短信格式,0是PDU模式，1是TEXT模式
	SetCMGF, err := machine.MsgSetCMGF("1")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(SetCMGF), Sub)
	}

	//获取当前短信格式,0是PDU模式，1是TEXT模式
	GetCMGF, err := machine.MsgGetCMGF()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(GetCMGF), Sub)
	}

	//设置短信URC输出的通道,分为："usbat"(USB的AT虚拟串口,默认),"usbmodem"(usb接入后，用于拨号上网的端口),"uart1"(硬件串口)
	QURCCFG, err := machine.MsgSetQURCCFG("usbat")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(QURCCFG), Sub)
	}

	//读取短信,("REC UNREAD","REC READ","STO UNSENT","STO SENT","ALL")
	GetCMGL, err := machine.MsgGetCMGL("ALL")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(GetCMGL), Sub)
	}

	//获取短信中心号码
	CSCA, err := machine.MsgGetCSCA()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CSCA), Sub)
	}

}
