package heap

type Func_Data struct {
	Args []string
	Line string
}

type Stack struct {
	BaseM map[string]*Heap
	FuncM map[string]*Heap
	FuncD map[string]*Func_Data
}

func (s *Stack) Override() {
	for name, heap := range s.BaseM {
		if heap.Global {
			s.FuncM[name] = heap
		}
	}
}

type VM struct {
	Files  map[string]*Stack
	Doing  string
	Import string
}
