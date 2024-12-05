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
	print("%v test : %v -> {&l} {test}")
	val b = "%v -> {&f}"
	print("%v -> {&b}")
	`
	vm.Get_Line("main", strings.Split(code, "\n"))
}
