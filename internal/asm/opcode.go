package asm

// TODO: reduce op size to byte to get a better
// memory footprint
type Opcode int

const (
	CONST Opcode = iota
	LOAD
	STORE
	ADD
	CALL
)

func (op Opcode) String() string {
	switch op {
	case CONST:
		return "CONST"
	case LOAD:
		return "LOAD"
	case STORE:
		return "STORE"
	case ADD:
		return "ADD"
	case CALL:
		return "CALL"
	default:
		panic("unreachable")
	}
}
