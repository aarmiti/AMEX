package main

import "github.com/aarmiti/AMEX/pkg/amex"

func main() {
	engine := amex.NewAmexEngine("microservice.yaml", "./output/")
	engine.SetupMicroService()

}
