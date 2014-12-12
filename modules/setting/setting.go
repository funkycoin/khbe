package setting

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var (
	RpcHost     string
	RpcUser     string
	RpcPassword string

	Cfg *goconfig.ConfigFile
)

func NewConfigContext() {
	var err error
	Cfg, err = goconfig.LoadConfigFile("config.ini")
	if err != nil {
		fmt.Printf("Fail to parse 'config.ini': %v\n", err)
	}
	RpcHost, err = Cfg.GetValue("", "rpchost")
	if err != nil {
		fmt.Printf("Unable to parse rpchost\n", err)
	}
	RpcUser, err = Cfg.GetValue("", "rpcuser")
	if err != nil {
		fmt.Printf("Unable to parse rpcuser", err)
	}
	RpcPassword, err = Cfg.GetValue("", "rpcpassword")
	if err != nil {
		fmt.Printf("Unable to parse rpcpassword", err)
	}
}
