package main

import (
	"strings"
	"vsce/vm"
)

func main() {
	code := `
	func test (x, y)
    {
	    var owo = "%v %v -> {&x} {&y}"

		print("%v -> {&owo}")
	}
	call test (12,23)
	print("%v %v -> {&x} {&y}")
	print("%v -> {&owo}")
	`
	vm.Get_Line("main", strings.Split(code, "\n"))
}
