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
	RETURN    interface{} = nil
)

var (
	VAR       bool        = false
	VAR_NALE  int         = 0
	VAR_TYPE  int         = 0
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
