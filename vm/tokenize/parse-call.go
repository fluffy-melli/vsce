package tokenize

import (
	"fmt"
	"strings"
	"vsce/vm/function"
	"vsce/vm/heap"
	"vsce/vm/heap/cash"
)

func Call_Parse(token TOKEN, file_lines []string, line string, i int, run bool) {
	switch token.Type {
	case LPAREN:
		cash.CALL_FUNC = strings.TrimSpace(line[cash.StartIndex : token.StartIndex+1])
		cash.CALL_FUNC = strings.ReplaceAll(cash.CALL_FUNC, "(", "")
		cash.CALL_PULL = true
		cash.StartIndex = token.StartIndex + 1
	case RPAREN:
		cash.CALL_PULL = false
		for _, v := range strings.Split(cash.CALL_BACK, ",") {
			txt := function.SprintF(fmt.Sprintf("%v", v), run)
			txt = strings.ReplaceAll(txt, "(", "")
			txt = strings.ReplaceAll(txt, ")", "")
			cash.CALL_ARGS = append(cash.CALL_ARGS, txt)
		}
		cash.Get_Stack().Override()
		funinfo := cash.Get_Stack().FuncD[cash.CALL_FUNC]
		for i, fi := range funinfo.Args {
			fi = strings.ReplaceAll(fi, "(", "")
			fi = strings.ReplaceAll(fi, ")", "")
			cash.Get_Stack().FuncM[strings.TrimSpace(fi)] = &heap.Heap{
				Const:  false,
				Global: false,
				Value:  cash.CALL_ARGS[i],
			}
		}
		cash.Clear_CALL()
		Inter(strings.Split(funinfo.Line, "\n"))
		cash.Get_Stack().Overwrite()
	default:
		cash.CALL_BACK += token.Literal
	}
}
