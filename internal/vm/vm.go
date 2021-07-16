package vm

import (
	"fmt"
	"strings"

	asm "github.com/hsiaosiyuan0/ripplet/internal/asm"
)

type Closure struct {
	fn     *asm.FnShape
	uppers map[string]interface{}
}

func NewClosure(fn *asm.FnShape) *Closure {
	return &Closure{
		fn:     fn,
		uppers: make(map[string]interface{}),
	}
}

type ObjType int

const (
	STR ObjType = iota
	NUM
	BOOL
	NIL
	DIC
	ARR
	CLOSURE
	GFN
)

type Object struct {
	Typ  ObjType
	Data interface{}
}

type Fn struct {
}

type Args struct {
	Len   int
	slots []*Object
}

var nilObj = &Object{
	Typ:  NIL,
	Data: nil,
}

func (a *Args) Get(i int) *Object {
	if i < a.Len {
		return a.slots[i]
	}
	return nilObj
}

type GoFnImpl func(*Args) interface{}

type GoFn struct {
	Name string
	Impl GoFnImpl
}

type CallFrame struct {
	*Closure
	pc int

	locals []*Object

	prev *CallFrame
}

func NewCallFrame(fn *asm.FnShape) *CallFrame {
	frame := &CallFrame{
		locals: make([]*Object, fn.LocalCnt),
	}
	frame.Closure = NewClosure(fn)
	return frame
}

type Vm struct {
	chunk     *asm.Chunk
	cf        *CallFrame
	opds      []*Object
	externals map[string]*Object
}

func NewVm(chunk *asm.Chunk) *Vm {
	return &Vm{
		chunk:     chunk,
		opds:      make([]*Object, 0),
		externals: make(map[string]*Object),
	}
}

func (v *Vm) pushOperand(obj *Object) {
	v.opds = append(v.opds, obj)
}

func (v *Vm) popOperand() *Object {
	cnt := len(v.opds)
	opd, opds := v.opds[cnt-1], v.opds[:cnt-1]
	v.opds = opds
	return opd
}

func (v *Vm) hasInstr() bool {
	return v.cf.pc != len(v.cf.fn.Instrs)
}

func (v *Vm) fetchInstr() int {
	op := v.cf.fn.Instrs[v.cf.pc]
	v.cf.pc += 1
	return op
}

func (v *Vm) fetchOp() asm.Opcode {
	instr := v.fetchInstr()
	return asm.Opcode(instr)
}

func (v *Vm) dispatch(op asm.Opcode) {
	switch op {
	case asm.CONST:
		c := v.chunk.GetConst(v.fetchInstr())
		typ := NUM
		if c.Typ == asm.ConstStr {
			typ = STR
		}
		v.pushOperand(&Object{Typ: typ, Data: c.Val})
	case asm.STORE:
		i := v.fetchInstr()
		v.cf.locals[i] = v.popOperand()
	case asm.LOAD_EXT:
		i := v.fetchInstr()
		nameConst := v.chunk.GetConst(i)
		name := nameConst.Val.(string)
		v.pushOperand(v.externals[name])
	case asm.CALL:
		argsLen := v.fetchInstr() - 1
		args := &Args{
			Len:   argsLen,
			slots: make([]*Object, argsLen),
		}
		for i := 0; i < argsLen; i++ {
			args.slots[i] = v.popOperand()
		}

		closure := v.popOperand()
		if closure.Typ == GFN {
			gfn := closure.Data.(GoFnImpl)
			gfn(args)
		} else {
			panic("not implemented")
		}
	case asm.CONCAT:
		subsLen := v.fetchInstr()
		subs := make([]string, subsLen)

		for i := 0; i < subsLen; i++ {
			opd := v.popOperand()
			// FIXME: toString
			ss := opd.Data.(string)
			subs = append(subs, ss)
		}

		v.pushOperand(&Object{
			Typ:  STR,
			Data: strings.Join(subs, ""),
		})
	case asm.LOAD:
		i := v.fetchInstr()
		v.pushOperand(v.cf.locals[i])
	case asm.CLOSURE:
		fmt.Printf(op.String() + "\n")
		fmt.Printf("INDEX_%d\n", v.fetchInstr())
	case asm.LOAD_UP, asm.STORE_UP:
		fmt.Printf(op.String() + "\n")
		fmt.Printf("NAME_%s\n", v.chunk.ConstStr(v.fetchInstr()))
	default:
		fmt.Printf("%s\n", op.String())
	}
}

func (v *Vm) AddExtFn(fn *GoFn) *Vm {
	v.externals[fn.Name] = &Object{
		Typ:  GFN,
		Data: fn.Impl,
	}
	return v
}

func (v *Vm) externalKeys() []string {
	keys := make([]string, 0, len(v.externals))
	for key := range v.externals {
		keys = append(keys, key)
	}
	return keys
}

func (v *Vm) SetCode(code string) *Vm {
	v.chunk = asm.Compile(code, v.externalKeys())
	return v
}

func (v *Vm) Exec() {
	v.cf = NewCallFrame(&v.chunk.Fn)

	for v.hasInstr() {
		op := v.fetchOp()
		v.dispatch(op)
	}
}
