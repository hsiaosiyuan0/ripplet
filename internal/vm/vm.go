package vm

import (
	"fmt"

	asm "github.com/hsiaosiyuan0/ripplet/internal/asm"
)

type Closure struct {
	Fn     *asm.FnShape
	Uppers map[string]interface{}
}

func NewClosure(fn *asm.FnShape) *Closure {
	return &Closure{
		Fn:     fn,
		Uppers: make(map[string]interface{}),
	}
}

type CallFrame struct {
	*Closure
	PC     int
	Locals map[string]interface{}
	Prev   *CallFrame
}

func NewCallFrame(fn *asm.FnShape) *CallFrame {
	frame := &CallFrame{}
	frame.Closure = NewClosure(fn)
	return frame
}

type Vm struct {
	cf *CallFrame
}

func NewVm() *Vm {
	return &Vm{}
}

func (s *Vm) hasInstr() bool {
	return s.cf.PC != len(s.cf.Fn.Instrs)
}

func (s *Vm) fetchInstr() int {
	op := s.cf.Fn.Instrs[s.cf.PC]
	s.cf.PC += 1
	return op
}

func (s *Vm) fetchOp() asm.Opcode {
	instr := s.fetchInstr()
	return asm.Opcode(instr)
}

func (s *Vm) dispatch(op asm.Opcode) {
	switch op {
	case asm.CONST, asm.STORE, asm.LOAD, asm.CALL:
		fmt.Printf("OP: %s Index: %d\n", op.String(), s.fetchInstr())
	case asm.ADD:
		fmt.Printf("OP: %s\n", op.String())
	}

}

func (s *Vm) Exec(fn *asm.FnShape) {
	s.cf = NewCallFrame(fn)

	for s.hasInstr() {
		op := s.fetchOp()
		s.dispatch(op)
	}
}
