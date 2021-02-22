package serialization

import "AT_GO/mode"

type MSGSerialization interface {
	MSGCreate(mode.SMS) ([]byte, error)
	//Dec([]byte)(Sms mode.SMS)
}
