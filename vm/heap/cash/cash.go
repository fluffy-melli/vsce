package cash

var (
	StartIndex int = 0
	EndIndex   int = 0
)

var (
	IF bool = false
)

var (
	CALL      bool        = false
	CALL_FUNC string      = ""
	CALL_PULL bool        = false
	CALL_BACK string      = ""
	CALL_ARGS []string    = make([]string, 0)
	CALL_PASS int         = 0
	RETURN    interface{} = nil
)

var (
	PRINT      bool   = false
	PRINT_OUT  string = ""
	PRINT_PULL bool   = false
)

var (
	VAR       bool        = false
	VAR_NALE  int         = 0
	VAR_TYPE  int         = 0
	VAR_PASS  int         = 0
	VAR_LONG  bool        = false
	VAR_NAME  string      = ""
	VAR_VALUE interface{} = nil
)

var (
	FUNC      bool     = false
	FUNC_NALE int      = 0
	FUNC_PASS int      = 0
	FUNC_PUSH bool     = false
	FUNC_PUSD int      = 0
	FUNC_NAME string   = ""
	FUNC_ARGS []string = make([]string, 0)
	FUNC_LINE string   = ""
)

func Clear_CALL() {
	CALL = false
	CALL_FUNC = ""
	CALL_PULL = false
	CALL_BACK = ""
	CALL_ARGS = make([]string, 0)
	RETURN = nil
}

func Clear_VAR() {
	VAR = false
	VAR_NALE = 0
	VAR_TYPE = 0
	VAR_PASS = 0
	VAR_LONG = false
	VAR_NAME = ""
	VAR_VALUE = nil
}

func Clear_FUNC() {
	FUNC = false
	FUNC_NALE = 0
	FUNC_PASS = 0
	FUNC_PUSH = false
	FUNC_PUSD = 0
	FUNC_NAME = ""
	FUNC_ARGS = make([]string, 0)
	FUNC_LINE = ""
}

func Clear_PRINT() {
	PRINT = false
	PRINT_OUT = ""
	PRINT_PULL = false
}
