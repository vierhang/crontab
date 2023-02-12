package main

import (
	"flag"
	"fmt"
	"github.com/vierhang/crontab/master"
)

var (
	configFile string
)

func initArgs() {
	flag.StringVar(&configFile, "config", "./master.json", "传入指定 master.json 路径")
	flag.Parse()
}

func main() {
	var (
		err error
	)
	initArgs()
	if err = master.InitConfig(configFile); err != nil {
		panic(err)
	}

	if err = master.InitJobMgr(); err != nil {
		panic(err)
	}

	if err = master.InitApiServer(); err != nil {
		panic(err)
	}

	fmt.Println("start success")
}
