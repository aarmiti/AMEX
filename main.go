package main

import "github.com/aarmiti/AMEX/pkg/amex"

func main() {
	p := amex.NewAmex("manifest.yaml", "./output/")
	p.SetupMicroService()

}
