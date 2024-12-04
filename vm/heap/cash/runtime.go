package cash

import "vsce/vm/heap"

var Runtime heap.VM = heap.VM{
	Files:  make(map[string]*heap.Stack),
	Doing:  "",
	Import: "",
}
