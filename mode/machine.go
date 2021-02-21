package mode

import (
	"github.com/tarm/serial"
	"time"
)

type Machine struct {
	MachineName string

	ComName     string
	Baud        int
	ReadTimeout time.Duration
	Size        byte
	Parity      serial.Parity
	StopBits    serial.StopBits
}
