package tokenize

import "vsce/vm/heap/cash"

func Inter(file_lines []string) {
	for i, line := range file_lines {
		tokens := GET_OPERATOR(line)
		if len(tokens) == 0 {
			continue
		}
		for _, token := range tokens {
			if !cash.VAR && !cash.IF && !cash.FUNC && !cash.CALL {
				switch token.Type {
				case FUNC:
					cash.FUNC = true
					cash.StartIndex = token.EndIndex + 1
					cash.FUNC_NALE = i
					continue
				case VAR, VAL, CONST:
					cash.VAR = true
					cash.VAR_TYPE = token.Type
					cash.VAR_NALE = i
					cash.StartIndex = token.EndIndex + 1
					continue
				case PRINT, PRINTF, PRINTLN:
					cash.PRINT = true
					cash.StartIndex = token.EndIndex + 1
					continue
				case CALL:
					cash.CALL = true
					cash.StartIndex = token.EndIndex + 1
					continue
				}
			}
			if cash.CALL {
				Call_Parse(token, file_lines, line, i, true)
			}
			if cash.FUNC {
				Func_Parse(token, file_lines, line, i)
			}
			if cash.VAR {
				Variable_Parse(token, file_lines, line, i, true)
			}
			if cash.PRINT {
				Print_Parse(token, file_lines, line, i, true)
			}
		}
	}
}
