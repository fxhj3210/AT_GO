package mode

import (
	"time"
)

type MSGInfo struct {
	Index string
	State string
	MSG
}

type MSG struct {
	Source int
	Body   []byte
	Time   time.Duration
}
