package main

import (
	"strings"
	"vsce/vm"
)

func main() {
	code := `
	val f = 123.12
	val l = [
	    1,2,3,4,5,[
		    "123"
		]
	]
	val j = {
	    "owo":123
	}
	`
	vm.Get_Line(strings.Split(code, "\n"))
}
