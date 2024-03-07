package main

import (
	"github.com/Elsharaky/Cars-API.git/app"
	"github.com/Elsharaky/Cars-API.git/config"
)

func main() {
	if err := app.SetupAndRunAPI(); err != nil {
		panic(err.Error())
	}

	defer config.GetDB().Close()
}
