package asm

import "fmt"

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
	LOAD_ARG
	ADD
	SUB
	MUL
	DIV
	MOD
	ARR
	SUBSCRIPT
	CONCAT
	CALL
	RET
	NOT
	TEST
	JMP_F
	JMP
	IS
	IS_NOT
	GT
	LT
	GE
	LE
	OR
	AND
	NEG
	BOOL_T
	BOOL_F
	NIL
)

func (op Opcode) String() string {
	switch op {
	case CONST:
		return "CONST"
	case LOAD:
		return "LOAD"
	case STORE:
		return "STORE"
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
	case LOAD_ARG:
		return "LOAD_ARG"
	case RET:
		return "RET"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case ARR:
		return "ARR"
	case SUBSCRIPT:
		return "SUBSCRIPT"
	case TEST:
		return "TEST"
	case JMP_F:
		return "JMP_F"
	case JMP:
		return "JMP"
	case IS:
		return "IS"
	case IS_NOT:
		return "IS_NOT"
	case GT:
		return "GT"
	case GE:
		return "GE"
	case LT:
		return "LT"
	case LE:
		return "LE"
	case OR:
		return "OR"
	case AND:
		return "AND"
	case NEG:
		return "NEG"
	case BOOL_T:
		return "BOOL_T"
	case BOOL_F:
		return "BOOL_F"
	case NIL:
		return "NIL"
	default:
		panic(fmt.Errorf("unreachable: %d", op))
	}
}
