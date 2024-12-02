package tokenize

import (
	"fmt"
	"strings"
)

func ParseVariableAssignment(tokens []Token, variables map[string]interface{}, inFunctionBody bool) error {
	var varName string
	var value interface{}
	for _, token := range tokens {
		switch token.Type {
		case TOKEN_VAR:
			if inFunctionBody {
				continue
			}
			varName = strings.TrimPrefix(token.Literal, "val ")
			if _, exists := variables[varName]; exists {
				return fmt.Errorf("error: variable %s already exists", varName)
			}
		case TOKEN_ASSIGN:
		case TOKEN_NUMBER, TOKEN_STRING:
			switch token.Type {
			case TOKEN_NUMBER:
				value = token.Literal
			case TOKEN_STRING:
				value = strings.Trim(token.Literal, `"`)
			}
			variables[varName] = value
			fmt.Printf("Variable %s assigned value: %v\n", varName, value)
			return nil
		}
	}
	return nil
}

// ParseFunctionDeclaration은 함수 선언과 본문을 처리합니다.
// 본문을 나중에 따로 처리할 수 있도록 본문을 반환
func ParseFunctionDeclaration(tokens []Token) (string, []string) {
	var funcName string
	var args []string
	var body []string
	var inArgs bool
	var inBody bool

	// 함수 이름과 인자 파싱
	for _, token := range tokens {
		switch token.Type {
		case TOKEN_FUNC:
			funcName = strings.TrimPrefix(token.Literal, "func ")
			index := strings.Index(funcName, "(")
			if index != -1 {
				funcName = funcName[:index]
				inArgs = true
			}
			funcName = strings.TrimSpace(funcName)
		case TOKEN_IDENTIFIER:
			if inArgs {
				args = append(args, token.Literal)
			}
		case TOKEN_PUNCTUATOR:
			if token.Literal == "(" {
				inArgs = true
			} else if token.Literal == ")" {
				inArgs = false
			} else if token.Literal == "," && inArgs {
				continue
			} else if token.Literal == "{" {
				inBody = true
			} else if token.Literal == "}" {
				inBody = false
			}
		}

		// 본문 파싱은 나중에 처리할 수 있도록 따로 저장
		if inBody && token.Literal != "{" && token.Literal != "}" {
			body = append(body, token.Literal)
		}
	}

	// 함수 이름과 본문 반환
	fmt.Printf("Function %s with arguments: %v\n", funcName, args)
	return funcName, body
}
