package main

import (
	"strings"
	"vsce/vm"
)

func main() {
	code := `
	func test 
	(x, y) 
	{
	    if 
		(x >= y) 
		{
		    return 0
		}
	}
	call test(0,2)
	`
	vm.Get_Line(strings.Split(code, "\n"))
}
