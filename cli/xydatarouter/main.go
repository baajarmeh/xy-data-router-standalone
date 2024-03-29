package main

import (
	"flag"
	"fmt"

	"github.com/fufuok/utils/xdaemon"

	"github.com/fufuok/xy-data-router/conf"
	"github.com/fufuok/xy-data-router/master"
)

var version bool

func init() {
	flag.StringVar(&conf.ConfigFile, "c", conf.ConfigFile, "配置文件绝对路径(可选)")
	flag.StringVar(&conf.ForwardTunnel, "f", "", "指定上联 Tunnel 地址(可选)")
	flag.BoolVar(&version, "v", false, "版本信息")
	flag.Parse()
}

func main() {
	if version {
		fmt.Println(">>>", conf.APPName, conf.Version, conf.GoVersion)
		fmt.Println(">>>", conf.GitCommit)
		return
	}

	conf.InitConfig()

	if !conf.Debug {
		xdaemon.NewDaemon(conf.LogDaemon).Run()
	}

	defer master.Stop()
	master.Start()
	master.Watcher()

	select {}
}
