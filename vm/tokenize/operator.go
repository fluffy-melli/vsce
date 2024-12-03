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
)

type TOKEN struct {
	Type    int
	Literal string
}

const pattern = `(?i)(?P<or>\|\|)|(?P<and>&&)|(?P<not>!)|(?P<equal>=)|(?P<du_equal>==)|(?P<not_equal>!=)|(?P<lt_equal><=)|(?P<gt_equal>>=)|(?P<lt><)|(?P<gt>>)|(?P<plus>\+)|(?P<minus>-)|(?P<multi>\*)|(?P<square>\*\*)|(?P<divide>/)|(?P<remain>%)|(?P<lparen>\()|(?P<rparen>\))|(?P<lpfunc>\{)|(?P<rpfunc>\})|(?P<comma>,)|(?P<var>\bvar\b)|(?P<val>\bval\b)|(?P<const>\bconst\b)|(?P<else_if>\belse if\b)|(?P<else>\belse\b)|(?P<if>\bif\b)|(?P<func>\bfunc\b)|(?P<fors>\bfor\b)|(?P<whiles>\bwhile\b)|(?P<break>\bbreak\b)|(?P<continue>\bcontinue\b)|(?P<return>\breturn\b)|(?P<print>\bprint\b)|(?P<printf>\bprintf\b)|(?P<println>\bprintln\b)|(?P<import>\bimport\b)|(?P<call>\bcall\b)|(?P<number>\d+)|(?P<string>"[^"]*")`

func GET_OPERATOR(input string) []TOKEN {
	var tokens []TOKEN
	input = strings.TrimSpace(input)
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		for i := range re.SubexpNames() {
			if i > 0 && match[i] != "" {
				var token = TOKEN{Type: i, Literal: match[i]}
				tokens = append(tokens, token)
			}
		}
	}

	return tokens
}
