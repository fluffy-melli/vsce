package main

import (
	"strings"
	"vsce/vm"
	"vsce/vm/function"
)

func main() {
	filename := "./main.vc"
	con := function.Read(filename)
	vm.Get_Line(filename, strings.Split(con, "\n"))
}
