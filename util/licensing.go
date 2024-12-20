package util

import (
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/config"
)

type Licensing interface {
	InitApp() error
}

type licensing struct {
	config config.CfgStruct
}

func NewLicensing(config config.CfgStruct) Licensing {
	return &licensing{
		config: config,
	}
}

func (l *licensing) InitApp() error {
	fmt.Println("\033[1;36m██████╗ ██████╗  ██████╗ ███████╗██╗██╗  ██╗ ██████╗ ██████╗ ███╗   ██╗\033[0m")
	fmt.Println("\033[1;36m██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║╚██╗██╔╝██╔════╝██╔═══██╗████╗  ██║\033[0m")
	fmt.Println("\033[1;36m██████╔╝██████╔╝██║   ██║█████╗  ██║ ╚███╔╝ ██║     ██║   ██║██╔██╗ ██║\033[0m")
	fmt.Println("\033[1;36m██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║ ██╔██╗ ██║     ██║   ██║██║╚██╗██║\033[0m")
	if l.config.App.Server.Port != "" {
		fmt.Println("\033[1;36m██║     ██║  ██║╚██████╔╝██║     ██║██╔╝ ██╗╚██████╗╚██████╔╝██║ ╚████║\033[0m	Service	:", "\033[1;32m", l.config.App.Name, "\033[0m")
		fmt.Println("\033[1;36m╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝\033[0m	Port	:", "\033[1;32m", l.config.App.Server.Port, "\033[0m")
	} else {
		fmt.Println("\033[1;36m██║     ██║  ██║╚██████╔╝██║     ██║██╔╝ ██╗╚██████╗╚██████╔╝██║ ╚████║\033[0m")
		fmt.Println("\033[1;36m╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝\033[0m	Service	:", "\033[1;32m", l.config.App.Name, "\033[0m")
	}

	return nil
}
