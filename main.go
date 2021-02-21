package main

import (
	"AT_GO/ec200s"
	"fmt"
)

const Sub = "\n----------------------------"

func main() {
	machine := ec200s.New("COM4")

	//设置TA响应模式,0时：0,1时：ok
	ATV, err := machine.SysSetATV("1")
	if err != nil {
		fmt.Println("ATV=", err)
	} else {
		fmt.Println("ATV=", ATV, Sub)
	}

	//显示产品标识信息(显示全部）
	ATI, err := machine.SysGetATI()
	if err != nil {
		fmt.Println("ATI=", err)
	} else {
		fmt.Println("ATI=", ATI, Sub)
	}

	//请求产品序列号标识
	CGSN, err := machine.SysGetCGSN()
	if err != nil {
		fmt.Println("CGSN=", err)
	} else {
		fmt.Println("CGSN=", CGSN, Sub)
	}

	//查询 SIM 卡状态，返回 READY 则表示SIM卡正常
	CPIN, err := machine.SysGetCPIN()
	if err != nil {
		fmt.Println("CPIN=", err)
	} else {
		fmt.Println("CPIN=", CPIN, Sub)
	}

	//查询模组是否注册上GSM网络.如果返回 1 或 5 ，代表 CS 服务注册成功。1 表示已注册上本地网,5表示注册上漫游网
	CREG, err := machine.NetworkGetCREG()
	if err != nil {
		fmt.Println("CREG=", err)
	} else {
		fmt.Println("CREG=", CREG, Sub)
	}

	//查询模组是否注册上GPRS网络，+CGREG:0,1 表示已注册上本地网，+CGREG:0,5表示注册上漫游网。
	CGREG, err := machine.NetworkGetCGREG()
	if err != nil {
		fmt.Println("CGREG=", err)
	} else {
		fmt.Println("CGREG=", CGREG, Sub)
	}

}
