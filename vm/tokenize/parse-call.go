package tokenize

import (
	"fmt"
	"strings"
	"vsce/vm/function"
	"vsce/vm/heap"
	"vsce/vm/heap/cash"
)

func Call_Parse(token TOKEN, file_lines []string, line string, i int) {
	switch token.Type {
	case LPAREN:
		if cash.CALL_PULL {
			cash.CALL_BACK += "("
			cash.CALL_PASS += 1
		} else {
			cash.CALL_FUNC = strings.TrimSpace(line[cash.StartIndex : token.StartIndex+1])
		}
		cash.CALL_PULL = true
		cash.StartIndex = token.StartIndex + 1
	case RPAREN:
		if cash.FUNC_PASS != 0 {
			cash.CALL_BACK += ")"
			cash.FUNC_PASS -= 1
			return
		}
		cash.CALL_PULL = false
		for _, v := range strings.Split(cash.CALL_BACK, ",") {
			cash.CALL_ARGS = append(cash.CALL_ARGS, function.SprintF(fmt.Sprintf("%v", v)))
		}
		cash.CALL_RUN = true
		cash.Get_Stack().Override()
		funinfo := cash.Get_Stack().FuncD[cash.CALL_FUNC]
		for i, fi := range funinfo.Args {
			cash.Get_Local_Stack()[strings.TrimSpace(fi)] = &heap.Heap{
				Const:  false,
				Global: false,
				Value:  cash.CALL_ARGS[i],
			}
		}
		cash.CALL = false
		Inter(strings.Split(funinfo.Line, "\n"))
		cash.Get_Stack().Overwrite()
		cash.Clear()
	default:
		cash.CALL_BACK += token.Literal
	}
}

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
				Call_Parse(token, file_lines, line, i)
			}
			if cash.FUNC {
				Func_Parse(token, file_lines, line, i)
			}
			if cash.VAR {
				Variable_Parse(token, file_lines, line, i)
			}
			if cash.PRINT {
				Print_Parse(token, file_lines, line, i)
			}
		}
	}
}
