package vm

import (
	"vsce/vm/cash"
	"vsce/vm/tokenize"
)

func Get_Line(file_lines []string) {
	for i, line := range file_lines {
		tokens := tokenize.GET_OPERATOR(line)
		if len(tokens) == 0 {
			continue
		}
		for _, token := range tokens {
			if !cash.VAR && !cash.IF && !cash.FUNC {
				switch token.Type {
				case tokenize.FUNC:
					cash.FUNC = true
					cash.StartIndex = token.EndIndex + 1
					cash.FUNC_NALE = i
					continue
				case tokenize.VAR, tokenize.VAL, tokenize.CONST:
					cash.VAR = true
					cash.VAR_TYPE = token.Type
					cash.VAR_NALE = i
					cash.StartIndex = token.EndIndex + 1
					continue
				}
			}
			if cash.FUNC {
				tokenize.Func_Parse(token, file_lines, line, i)
			}
			if cash.VAR {
				tokenize.Variable_Parse(token, file_lines, line, i)
			}
		}
	}
}
