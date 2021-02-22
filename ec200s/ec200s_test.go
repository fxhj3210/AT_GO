package ec200s

import (
	"fmt"
	"testing"
)

const Sub = "\n----------------------------"

func Test(t *testing.T) {
	machine := New("COM12")

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

	//查询错误讯息格式,1启用结果代码并使用数值(ERROR:10);2启用结果代码并使用详细值(ERROR:SIM not inserted)
	GetCMEE, err := machine.SysGetCMEE()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(GetCMEE), Sub)
	}

	//设置错误讯息格式,0禁用结果代码;1启用结果代码并使用数值;2启用结果代码并使用详细值
	SetCMEE, err := machine.SysSetCMEE("2")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(SetCMEE), Sub)
	}

	//查询 SIM 卡状态，返回 READY 则表示SIM卡正常
	CPIN, err := machine.SysGetCPIN()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CPIN), Sub)
	}

	//查询当前SIM卡手机号
	CNUM, err := machine.SysGetCNUM()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CNUM), Sub)
	}

	//获取国际移动用户识别码
	CIMI, err := machine.SysGetCIMI()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CIMI), Sub)
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

	//查询模组是否注册上 NB 网络，0 未注册，MT 当前当前没有搜索或者注册到运营商网络;1 已注册，注册到归属网络;2 未注册;3 注册被拒绝;4 未知错误;5 已注册，注册到漫游网络
	CEREG, err := machine.NetworkGetCEREG()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CEREG), Sub)
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

	//读取短信列表,("REC UNREAD","REC READ","STO UNSENT","STO SENT","ALL")
	//(“REC UNREAD”-未读;“REC READ”-已读;“STO UNSENT”-待发;“STO SENT”-已发;“ALL”-全部)
	GetCMGL, err := machine.MsgGetCMGL("ALL")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(GetCMGL), Sub)
	}

	//读取单条短信
	GetCMGR, err := machine.MsgGetCMGR("1")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(GetCMGR), Sub)
	}

	//删除短信
	/*
		DelCMGD, err := machine.MsgDelCMGD("0","")
		if err != nil {
			t.Errorf(err.Error())
		} else {
			fmt.Println(string(DelCMGD), Sub)
		}
	*/

	//获取短信中心号码
	CSCA, err := machine.MsgGetCSCA()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(string(CSCA), Sub)
	}

	/*
		//重启模块,会断电
		Restart, err := machine.SysRestart()
		if err != nil {
			t.Errorf(err.Error())
		} else {
			fmt.Println(string(Restart), Sub)
		}
	*/

	//编辑短信

}
