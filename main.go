package main

import (
	"log"

	"github.com/test-web/controller"
)

func main() {
	log.Println("started APP")
	c := controller.NewController()
	c.RunController()
}
