package main

import (
	"github.com/nicholasjackson/smi-controller/mesh"
	"github.com/nicholasjackson/smi-controller/pkg"
)

func main() {
	// register our lifecycle callbacks with the controller
	mesh.API().RegisterV1Alpha2(&loggerV2{})

	// create and start a the controller
	c := &pkg.SMIController{}
	c.Start()
}
