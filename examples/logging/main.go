package main

import (
	"github.com/nicholasjackson/smi-controller/sdk"
	"github.com/nicholasjackson/smi-controller/sdk/controller"
)

func main() {
	// register our lifecycle callbacks with the controller
	sdk.API().RegisterV1Alpha2(&loggerV2{})

	// create and start a the controller
	controller.Start()
}
