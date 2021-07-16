package asm

// TODO: reduce op size to byte to get a better
// memory footprint
type Opcode int

const (
	CONST Opcode = iota
	LOAD
	LOAD_UP
	STORE
	STORE_UP
	LOAD_EXT
	CLOSURE
	ADD
	CONCAT
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
	case CLOSURE:
		return "CLOSURE"
	case CONCAT:
		return "CONCAT"
	case LOAD_UP:
		return "LOAD_UP"
	case STORE_UP:
		return "STORE_UP"
	case LOAD_EXT:
		return "LOAD_EXT"
	default:
		panic("unreachable")
	}
}
