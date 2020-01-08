package main

import (
	"github.com/akbarpambudi/shoppies/injector"
)

func main() {
	shoppiesRouter, err := injector.InitializeShoppiesRouter()
	if err != nil {
		panic(err)
	}
	shoppiesRouter.BindProductDelivery(injector.InitializeProductDelivery())
	shoppiesRouter.Serve()
}
