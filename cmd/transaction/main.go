package main

import (
	_ "github.com/jsperandio/transaction/docs"

	fxserver "github.com/jsperandio/transaction/app/fx/module/server"
)

func main() {
	fxserver.Start()
}
