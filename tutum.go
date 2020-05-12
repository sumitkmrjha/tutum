package main

//go:generate go run directives_generate.go
//go:generate go run owners_generate.go

import (
	"github.com/sumitkmrjha/tutum/src/core"
)

func main() {
	coremain.Run()
}