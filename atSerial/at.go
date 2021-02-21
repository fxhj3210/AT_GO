package atSerial

type AT interface {
	MsgSetQURCCFG(string) ([]byte, error)
	NetworkGetCGREG() ([]byte, error)
	NetworkGetCREG() ([]byte, error)
	SysGetATI() ([]byte, error)
	SysGetCGSN() ([]byte, error)
	SysGetCPIN() ([]byte, error)
	SysSetATV(string) ([]byte, error)
	MsgGetCSCA() ([]byte, error)
	MsgGetCPMS() ([]byte, error)
	MsgGetCPMSState() ([]byte, error)
	MsgSetCPMS(string, string, string) ([]byte, error)
	MsgGetCMGF() ([]byte, error)
	MsgSetCMGF(string) ([]byte, error)
	MsgGetCMGL(string) ([]byte, error)
	NetworkGetCOPS() ([]byte, error)
}
