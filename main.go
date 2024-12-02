package main

import (
	"fmt"
	"vsce/vm/tokenize"
)

func main() {
	code := `
	val test = "1te3"
	func add (x, y) {
	    val test2 = x + y
	    return test
	}`

	tokens, err := tokenize.Tokenize(code)
	if err != nil {
		fmt.Println("Error tokenizing input:", err)
		return
	}

	variables := make(map[string]interface{})
	funcBodies := make(map[string][]string) // 함수 이름에 대응하는 본문을 저장할 맵

	// 파서 처리
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == tokenize.TOKEN_VAR {
			err := tokenize.ParseVariableAssignment(tokens[i:], variables, false)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if tokens[i].Type == tokenize.TOKEN_FUNC {
			funcName, body := tokenize.ParseFunctionDeclaration(tokens[i:])
			funcBodies[funcName] = body
		}
	}

	fmt.Println(variables)
	fmt.Println(funcBodies)
}
