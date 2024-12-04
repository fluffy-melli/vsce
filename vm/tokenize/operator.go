package tokenize

import (
	"regexp"
	"strings"
)

const (
	// operator
	UNKNOW    = iota // unknow
	OR               // ||
	AND              // &&
	NOT              // !
	EQUAL            // =
	DU_EQUAL         // ==
	NOT_EQUAL        // !=
	LT_EQUAL         // <=
	GT_EQUAL         // >=
	LT               // <
	GT               // >
	PLUS             // +
	MINUS            // -
	MULTI            // *
	SQUARE           // **
	DIVIDE           // /
	REMAIN           // %
	LPAREN           // (
	RPAREN           // )
	LPFUNC           // {
	RPFUNC           // }
	COMMA            // ,
	//////////////////////////////
	// variable
	VAR   // var
	VAL   // val
	CONST // const
	//////////////////////////////
	// function
	ELSE_IF  // else if
	ELSE     // else
	IF       // if
	FUNC     // func
	FOR      // for
	WHILE    // while
	BREAK    // break
	CONTINUE // continue
	RETURN   // return
	PRINT    // print
	PRINTF   // printf
	PRINTLN  // println
	IMPORT   // import
	CALL     // call
	//////////////////////////////
	// value
	NUMBER
	STRING
	OLIST // [
	CLIST // ]
)

type TOKEN struct {
	Type       int
	Literal    string
	StartIndex int
	EndIndex   int
}

const pattern = `(?i)(?P<or>\|\|)|(?P<and>&&)|(?P<not>!)|(?P<equal>=)|(?P<du_equal>==)|(?P<not_equal>!=)|(?P<lt_equal><=)|(?P<gt_equal>>=)|(?P<lt><)|(?P<gt>>)|(?P<plus>\+)|(?P<minus>-)|(?P<multi>\*)|(?P<square>\*\*)|(?P<divide>/)|(?P<remain>%)|(?P<lparen>\()|(?P<rparen>\))|(?P<lpfunc>\{)|(?P<rpfunc>\})|(?P<comma>,)|(?P<var>\bvar\b)|(?P<val>\bval\b)|(?P<const>\bconst\b)|(?P<else_if>\belse if\b)|(?P<else>\belse\b)|(?P<if>\bif\b)|(?P<func>\bfunc\b)|(?P<fors>\bfor\b)|(?P<whiles>\bwhile\b)|(?P<break>\bbreak\b)|(?P<continue>\bcontinue\b)|(?P<return>\breturn\b)|(?P<print>\bprint\b)|(?P<printf>\bprintf\b)|(?P<println>\bprintln\b)|(?P<import>\bimport\b)|(?P<call>\bcall\b)|(?P<number>\d+)|(?P<string>"[^"]*")|(?P<olist>\[)|(?P<clist>\])`

func GET_OPERATOR(input string) []TOKEN {
	var tokens []TOKEN
	input = strings.TrimSpace(input)
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)
	indexMatches := re.FindAllStringSubmatchIndex(input, -1)
	for idx, match := range matches {
		startIndex := indexMatches[idx][0]
		endIndex := indexMatches[idx][1]
		for i, name := range re.SubexpNames() {
			if i > 0 && match[i] != "" {
				var token TOKEN
				switch name {
				case "or":
					token = TOKEN{Type: OR, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "and":
					token = TOKEN{Type: AND, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "not":
					token = TOKEN{Type: NOT, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "equal":
					token = TOKEN{Type: EQUAL, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "du_equal":
					token = TOKEN{Type: DU_EQUAL, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "not_equal":
					token = TOKEN{Type: NOT_EQUAL, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "lt_equal":
					token = TOKEN{Type: LT_EQUAL, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "gt_equal":
					token = TOKEN{Type: GT_EQUAL, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "lt":
					token = TOKEN{Type: LT, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "gt":
					token = TOKEN{Type: GT, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "plus":
					token = TOKEN{Type: PLUS, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "minus":
					token = TOKEN{Type: MINUS, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "multi":
					token = TOKEN{Type: MULTI, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "square":
					token = TOKEN{Type: SQUARE, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "divide":
					token = TOKEN{Type: DIVIDE, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "remain":
					token = TOKEN{Type: REMAIN, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "lparen":
					token = TOKEN{Type: LPAREN, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "rparen":
					token = TOKEN{Type: RPAREN, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "lpfunc":
					token = TOKEN{Type: LPFUNC, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "rpfunc":
					token = TOKEN{Type: RPFUNC, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "comma":
					token = TOKEN{Type: COMMA, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "var":
					token = TOKEN{Type: VAR, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "val":
					token = TOKEN{Type: VAL, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "const":
					token = TOKEN{Type: CONST, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "else_if":
					token = TOKEN{Type: ELSE_IF, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "else":
					token = TOKEN{Type: ELSE, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "if":
					token = TOKEN{Type: IF, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "func":
					token = TOKEN{Type: FUNC, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "fors":
					token = TOKEN{Type: FOR, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "whiles":
					token = TOKEN{Type: WHILE, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "break":
					token = TOKEN{Type: BREAK, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "continue":
					token = TOKEN{Type: CONTINUE, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "return":
					token = TOKEN{Type: RETURN, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "print":
					token = TOKEN{Type: PRINT, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "printf":
					token = TOKEN{Type: PRINTF, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "println":
					token = TOKEN{Type: PRINTLN, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "import":
					token = TOKEN{Type: IMPORT, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "call":
					token = TOKEN{Type: CALL, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "number":
					token = TOKEN{Type: NUMBER, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "string":
					token = TOKEN{Type: STRING, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "olist":
					token = TOKEN{Type: OLIST, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				case "clist":
					token = TOKEN{Type: CLIST, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				default:
					token = TOKEN{Type: UNKNOW, Literal: match[i], StartIndex: startIndex, EndIndex: endIndex}
				}
				tokens = append(tokens, token)
			}
		}
	}
	return tokens
}
