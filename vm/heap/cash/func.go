package cash

import "vsce/vm/heap"

func Get_Local_Stack() map[string]*heap.Heap {
	if CALL {
		return Runtime.Files[Runtime.Doing].FuncM
	}
	return Runtime.Files[Runtime.Doing].BaseM
}
