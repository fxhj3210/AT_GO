package mode

import "time"

/*
----------------------------------------
|	SCA		|
|服务中心号码	|
-----------------------------------
*/
type PDU struct {
	Source int
	Body   []byte
	Time   time.Duration
}

//TODO:PDU格式
