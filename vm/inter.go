package vm

import (
	"vsce/vm/heap"
	"vsce/vm/heap/cash"
	"vsce/vm/tokenize"
)

func Get_Line(file_name string, file_lines []string) {
	cash.Runtime.Doing = file_name
	if cash.Runtime.Files[cash.Runtime.Doing] == nil {
		cash.Runtime.Files[cash.Runtime.Doing] = &heap.Stack{
			BaseM: make(map[string]*heap.Heap),
			FuncM: make(map[string]*heap.Heap),
			FuncD: make(map[string]*heap.Func_Data),
		}
	}
	for i, line := range file_lines {
		tokens := tokenize.GET_OPERATOR(line)
		if len(tokens) == 0 {
			continue
		}
		for _, token := range tokens {
			if !cash.VAR && !cash.IF && !cash.FUNC && !cash.CALL {
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
				case tokenize.PRINT, tokenize.PRINTF, tokenize.PRINTLN:
					cash.PRINT = true
					cash.StartIndex = token.EndIndex + 1
					continue
				case tokenize.CALL:
					cash.CALL = true
					cash.StartIndex = token.EndIndex + 1
					continue
				}
			}
			if cash.CALL {
				tokenize.Call_Parse(token, file_lines, line, i)
			}
			if cash.FUNC {
				tokenize.Func_Parse(token, file_lines, line, i)
			}
			if cash.VAR {
				tokenize.Variable_Parse(token, file_lines, line, i)
			}
			if cash.PRINT {
				tokenize.Print_Parse(token, file_lines, line, i)
			}
		}
	}
}
