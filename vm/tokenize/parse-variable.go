package tokenize

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"vsce/vm/function"
	"vsce/vm/heap"
	"vsce/vm/heap/cash"
)

func parseJSON(input string) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Variable_Parse(token TOKEN, file_lines []string, line string, i int, run bool) {
	switch token.Type {
	case EQUAL:
		if cash.VAR_NALE != i {
			cash.VAR_NAME = strings.TrimSpace(file_lines[cash.FUNC_NALE][cash.StartIndex:])
			return
		}
		cash.VAR_NAME = strings.TrimSpace(line[cash.StartIndex:token.StartIndex])
		return
	case NUMBER:
		if cash.VAR_LONG {
			cash.VAR_VALUE = cash.VAR_VALUE.(string) + token.Literal
			return
		}
		value := strings.ReplaceAll(token.Literal, "_", "")
		n, e := strconv.Atoi(value)
		if e != nil {
			fmt.Println(e.Error())
			break
		}
		cash.VAR_VALUE = n
		cash.Get_Local_Stack(run)[cash.VAR_NAME] = &heap.Heap{
			Const:  cash.VAR_TYPE == CONST,
			Global: cash.VAR_TYPE == VAR,
			Value:  cash.VAR_VALUE,
		}
		cash.Clear_VAR()
		return
	case FLOAT:
		if cash.VAR_LONG {
			cash.VAR_VALUE = cash.VAR_VALUE.(string) + token.Literal
			return
		}
		value := strings.ReplaceAll(token.Literal, "_", "")
		n, e := strconv.ParseFloat(value, 64)
		if e != nil {
			fmt.Println(e.Error())
			break
		}
		cash.VAR_VALUE = n
		cash.Get_Local_Stack(run)[cash.VAR_NAME] = &heap.Heap{
			Const:  cash.VAR_TYPE == CONST,
			Global: cash.VAR_TYPE == VAR,
			Value:  cash.VAR_VALUE,
		}
		cash.Clear_VAR()
		return
	case STRING:
		if cash.VAR_LONG {
			cash.VAR_VALUE = cash.VAR_VALUE.(string) + token.Literal
			return
		}
		cash.VAR_VALUE = token.Literal
		cash.Get_Local_Stack(run)[cash.VAR_NAME] = &heap.Heap{
			Const:  cash.VAR_TYPE == CONST,
			Global: cash.VAR_TYPE == VAR,
			Value:  strings.ReplaceAll(function.SprintF(fmt.Sprintf("%v", cash.VAR_VALUE), run), "\"", ""),
		}
		cash.Clear_VAR()
		return
	case LPFUNC, OLIST: // JSON || LIST
		if cash.VAR_LONG {
			cash.VAR_VALUE = cash.VAR_VALUE.(string) + token.Literal
			cash.VAR_PASS += 1
			return
		}
		cash.VAR_LONG = true
		cash.VAR_VALUE = strings.TrimSpace(token.Literal)
		return
	case RPFUNC, CLIST:
		if cash.VAR_LONG && cash.VAR_PASS != 0 {
			cash.VAR_VALUE = cash.VAR_VALUE.(string) + token.Literal
			cash.VAR_PASS -= 1
			return
		}
		cash.VAR_LONG = false
		cash.VAR_VALUE = cash.VAR_VALUE.(string) + strings.TrimSpace(token.Literal)
		parsed, e := parseJSON(cash.VAR_VALUE.(string))
		if e != nil {
			fmt.Println(e.Error())
			break
		}
		cash.VAR_VALUE = parsed
		cash.Get_Local_Stack(run)[cash.VAR_NAME] = &heap.Heap{
			Const:  cash.VAR_TYPE == CONST,
			Global: cash.VAR_TYPE == VAR,
			Value:  cash.VAR_VALUE,
		}
		cash.Clear_VAR()
		return
	default:
		if cash.VAR_LONG {
			cash.VAR_VALUE = cash.VAR_VALUE.(string) + token.Literal
		}
		return
	}
}
