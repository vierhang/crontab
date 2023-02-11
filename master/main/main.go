package main

import (
	"github.com/vierhang/crontab/master"
)

func main() {
	var (
		err error
	)
	if err = master.InitApiServer(); err != nil {
		panic(err)
	}

}
