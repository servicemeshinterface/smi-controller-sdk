package main

import (
	"github.com/servicemeshinterface/smi-controller-sdk/examples/logging"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk/controller"
)

func main() {
	// register our lifecycle callbacks with the controller
	sdk.API().RegisterV1Alpha(&logging.Logger{})

	// create and start a the controller
	config := controller.DefaultConfig()
	controller.Start(config)
}
