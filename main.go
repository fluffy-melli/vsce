package main

import (
	"fmt"
	"vsce/vm/tokenize"
)

func main() {
	code := `
	func test (x, y) {
	    if (x >= y) {
		    return 0
		} else if (x == y) {
		    return 1
		} else {
		    return 2 
		}
	}
	call test(0,2)
	`

	tokens := tokenize.GET_OPERATOR(code)
	fmt.Println(tokens)
}
