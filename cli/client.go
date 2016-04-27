package cli

import flag "github.com/docker/docker/pkg/mflag"

// ClientFlags represents flags for the docker client.
type ClientFlags struct {
	FlagSet   *flag.FlagSet //预定义标志
	Common    *CommonFlags  //client和daemon的共同标志
	PostParse func()        //这是一个匿名函数,参见docker/docker/client.go

	ConfigDir string //配置文件目录.默认是$DOCKER_CONFIG,否则是$HOME/.docker
}
