package tokenize

import (
	"vsce/vm/heap/cash"
)

func Print_Parse(token TOKEN, file_lines []string, line string, i int) {
	switch token.Type {
	case LPAREN:
		cash.StartIndex = token.EndIndex + 1
		cash.PRINT_PULL = true
	case RPAREN:
		text := line[cash.StartIndex:token.StartIndex]
		GET_VAR(text)
		//fmt.Println(cash.PRINT_OUT)
		cash.Clear()
	case NUMBER, STRING:
		if !cash.PRINT_PULL {
			return
		}
		cash.PRINT_OUT += token.Literal
	case COMMA:
		if !cash.PRINT_PULL {
			return
		}
		cash.PRINT_OUT += " "
	}
}
