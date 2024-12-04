package tokenize

import (
	"strings"
	"vsce/vm/heap"
	"vsce/vm/heap/cash"
)

func Func_Parse(token TOKEN, file_lines []string, line string, i int) {
	switch token.Type {
	case LPAREN:
		if cash.FUNC_PUSH {
			return
		}
		if cash.FUNC_NALE != i {
			cash.FUNC_NAME = strings.TrimSpace(file_lines[cash.FUNC_NALE][cash.StartIndex+1:])
			cash.StartIndex = token.StartIndex + 1
			return
		}
		cash.FUNC_NAME = strings.TrimSpace(line[cash.StartIndex+1 : token.StartIndex+1])
		cash.StartIndex = token.StartIndex + 1
		return
	case RPAREN:
		if cash.FUNC_PUSH {
			return
		}
		cash.FUNC_ARGS = strings.Split((strings.TrimSpace(line[cash.StartIndex+1 : token.StartIndex+1])), ",")
		return
	case LPFUNC:
		if cash.FUNC_PUSH {
			cash.FUNC_LINE += "{\n"
			cash.FUNC_PASS += 1
			return
		}
		cash.FUNC_PUSH = true
		cash.FUNC_LINE += strings.TrimSpace(line[token.StartIndex+2:])
		return
	case RPFUNC:
		if cash.FUNC_PASS != 0 {
			cash.FUNC_LINE += "}\n"
			cash.FUNC_PASS -= 1
			return
		}
		cash.FUNC = false
		cash.FUNC_PUSH = false
		cash.FUNC_LINE += strings.TrimSpace(line[:token.StartIndex])
		cash.Runtime.Files[cash.Runtime.Doing].FuncD[cash.FUNC_NAME] = &heap.Func_Data{
			Args: cash.FUNC_ARGS,
			Line: cash.FUNC_LINE,
		}
		cash.Clear()
		return
	default:
		if cash.FUNC_PUSH && cash.FUNC_PUSD != i {
			cash.FUNC_LINE += strings.TrimSpace(line) + "\n"
			cash.FUNC_PUSD = i
		}
		return
	}
}
