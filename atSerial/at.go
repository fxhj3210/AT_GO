package atSerial

type AT interface {
	SysGetATI() (CGSN string, err error)
	SysGetCGSN() (CGSN string, err error)
	SysGetCPIN() (CGSN string, err error)
	SysSetATV(Model string) (CGSN string, err error)
	NetworkGetCGREG() (CGSN string, err error)
	NetworkGetCREG() (CGSN string, err error)
}
