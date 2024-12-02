package tokenize

import (
	"regexp"
	"strings"
)

// 토큰 타입 정의
const (
	TOKEN_VAR        = iota // 변수 선언 (val)
	TOKEN_FUNC              // 함수 선언 (func)
	TOKEN_IDENTIFIER        // 식별자 (변수명, 함수명)
	TOKEN_ASSIGN            // 할당 연산자 (=)
	TOKEN_NUMBER            // 숫자 (12)
	TOKEN_STRING            // 문자열 ("hello")
	TOKEN_PUNCTUATOR        // 구분자 (괄호, 중괄호)
	TOKEN_COMMA             // 쉼표 (인자 구분)
	TOKEN_UNKNOWN           // 알 수 없는 토큰
	TOKEN_OPERATOR          // 연산자 (+, -, *, /, ==, !=, <, > 등)
)

type Token struct {
	Type    int
	Literal string
}

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	input = strings.TrimSpace(input)

	// 연산자 패턴 추가: 산술 연산자, 비교 연산자
	pattern := `(?P<var>val\s+\w+)|(?P<func>func\s+\w+\s*\()|(?P<assign>=)|(?P<number>\d+)|(?P<string>"[^"]*")|(?P<identifier>\w+)|(?P<punctuator>[{}(),])|(?P<comma>,)|(?P<operator>\+|-|\*|/|==|!=|<|>|<=|>=)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		for i, name := range re.SubexpNames() {
			if i > 0 && match[i] != "" {
				var token Token
				switch name {
				case "var":
					token = Token{Type: TOKEN_VAR, Literal: match[i]}
				case "func":
					token = Token{Type: TOKEN_FUNC, Literal: match[i]}
				case "assign":
					token = Token{Type: TOKEN_ASSIGN, Literal: match[i]}
				case "number":
					token = Token{Type: TOKEN_NUMBER, Literal: match[i]}
				case "string":
					token = Token{Type: TOKEN_STRING, Literal: match[i]}
				case "identifier":
					token = Token{Type: TOKEN_IDENTIFIER, Literal: match[i]}
				case "punctuator":
					token = Token{Type: TOKEN_PUNCTUATOR, Literal: match[i]}
				case "comma":
					token = Token{Type: TOKEN_COMMA, Literal: match[i]}
				case "operator":
					token = Token{Type: TOKEN_OPERATOR, Literal: match[i]}
				}
				tokens = append(tokens, token)
			}
		}
	}

	return tokens, nil
}
