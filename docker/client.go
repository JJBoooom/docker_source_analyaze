package main

import (
	"path/filepath"

	"github.com/docker/docker/cli"
	"github.com/docker/docker/cliconfig"
	flag "github.com/docker/docker/pkg/mflag"
	"github.com/docker/docker/utils"
)

//创建客户端标志集
var clientFlags = &cli.ClientFlags{FlagSet: new(flag.FlagSet), Common: commonFlags}

//初始化clientFlags
func init() {
	//
	client := clientFlags.FlagSet
	//设置客户端的配置目录?
	client.StringVar(&clientFlags.ConfigDir, []string{"-config"}, cliconfig.ConfigDir(), "Location of client config files")

	//初始化PostParse方法
	clientFlags.PostParse = func() {
		clientFlags.Common.PostParse()

		if clientFlags.ConfigDir != "" {
			cliconfig.SetConfigDir(clientFlags.ConfigDir)
		}

		if clientFlags.Common.TrustKey == "" {
			clientFlags.Common.TrustKey = filepath.Join(cliconfig.ConfigDir(), defaultTrustKeyFile)
		}

		if clientFlags.Common.Debug {
			utils.EnableDebug()
		}
	}
}
