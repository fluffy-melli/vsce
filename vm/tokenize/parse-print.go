package tokenize

import (
	"fmt"
	"strings"
	"vsce/vm/function"
	"vsce/vm/heap/cash"
)

func Print_Parse(token TOKEN, file_lines []string, line string, i int) {
	switch token.Type {
	case LPAREN:
		cash.StartIndex = token.EndIndex + 1
		cash.PRINT_PULL = true
	case RPAREN:
		text := strings.ReplaceAll(line[cash.StartIndex:token.StartIndex], "\"", "")
		fmt.Println(function.SprintF(strings.ReplaceAll(text, "\"", "")))
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
