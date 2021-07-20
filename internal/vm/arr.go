package vm

type Array struct {
	cnt   int
	slots map[int]*Object
}

func NewArray(cnt int) *Array {
	return &Array{
		cnt:   cnt,
		slots: make(map[int]*Object),
	}
}

func (a *Array) Get(i int) *Object {
	if slot, ok := a.slots[i]; ok {
		return slot
	}
	a.slots[i] = nilObj
	return nilObj
}

func (a *Array) Set(i int, obj *Object) {
	a.slots[i] = obj
}
