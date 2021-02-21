package ec200s

import (
	"AT_GO/mode"
)

type ec200s struct {
	mode.Machine
}

func New(ComName string) *ec200s {
	return &ec200s{mode.Machine{
		MachineName: "EC200S",

		ComName:     ComName,
		Baud:        115200,
		ReadTimeout: 10,
		Size:        8,
		Parity:      0,
		StopBits:    1,
	}}
}
