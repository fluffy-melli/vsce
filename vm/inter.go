package vm

import (
	"fmt"
	"strconv"
	"strings"
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
				switch token.Type {
				case tokenize.LPAREN:
					if cash.FUNC_PUSH {
						continue
					}
					if cash.FUNC_NALE != i {
						cash.FUNC_NAME = strings.TrimSpace(file_lines[cash.FUNC_NALE][cash.StartIndex+1:])
						cash.StartIndex = token.StartIndex + 1
						continue
					}
					cash.FUNC_NAME = strings.TrimSpace(line[cash.StartIndex+1 : token.StartIndex+1])
					cash.StartIndex = token.StartIndex + 1
					continue
				case tokenize.RPAREN:
					if cash.FUNC_PUSH {
						continue
					}
					cash.FUNC_ARGS = strings.Split((strings.TrimSpace(line[cash.StartIndex+1 : token.StartIndex+1])), ",")
					continue
				case tokenize.LPFUNC:
					if cash.FUNC_PUSH {
						cash.FUNC_LINE += "{\n"
						cash.FUNC_PASS += 1
						continue
					}
					cash.FUNC_PUSH = true
					cash.FUNC_LINE += strings.TrimSpace(line[token.StartIndex+2:])
					continue
				case tokenize.RPFUNC:
					if cash.FUNC_PASS != 0 {
						cash.FUNC_LINE += "}\n"
						cash.FUNC_PASS -= 1
						continue
					}
					cash.FUNC = false
					cash.FUNC_PUSH = false
					cash.FUNC_LINE += strings.TrimSpace(line[:token.StartIndex])
					fmt.Println("New_Func :", cash.FUNC_NAME, cash.FUNC_ARGS, "\n", cash.FUNC_LINE)
					continue
				default:
					if cash.FUNC_PUSH && cash.FUNC_PUSD != i {
						cash.FUNC_LINE += strings.TrimSpace(line) + "\n"
						cash.FUNC_PUSD = i
					}
					continue
				}
			}
			if cash.VAR {
				switch token.Type {
				case tokenize.EQUAL:
					if cash.VAR_NALE != i {
						cash.FUNC_NAME = strings.TrimSpace(file_lines[cash.FUNC_NALE][cash.StartIndex:])
						continue
					}
					cash.FUNC_NAME = strings.TrimSpace(line[cash.StartIndex:token.StartIndex])
					continue
				case tokenize.NUMBER:
					value := strings.ReplaceAll(token.Literal, "_", "")
					n, e := strconv.Atoi(value)
					if e != nil {
						fmt.Println(e.Error())
						break
					}
					cash.VAR_VALUE = n
				case tokenize.FLOAT:
					value := strings.ReplaceAll(token.Literal, "_", "")
					n, e := strconv.ParseFloat(value, 64)
					if e != nil {
						fmt.Println(e.Error())
						break
					}
					cash.VAR_VALUE = n
				case tokenize.STRING:
					cash.VAR_VALUE = token.Literal
				case tokenize.LPFUNC, tokenize.OLIST: // JSON || LIST
					cash.VAR_LONG = true
				default:
					if cash.VAR_LONG {
					}
				}
			}
		}
	}
}
