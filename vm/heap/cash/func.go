package cash

import (
	"vsce/vm/heap"
)

func Get_Local_Stack(fun bool) map[string]*heap.Heap {
	if fun {
		return Runtime.Files[Runtime.Doing].FuncM
	}
	return Runtime.Files[Runtime.Doing].BaseM
}

func Get_Stack() *heap.Stack {
	return Runtime.Files[Runtime.Doing]
}
