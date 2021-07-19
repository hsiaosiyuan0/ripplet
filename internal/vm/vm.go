package vm

import (
	"fmt"
	"strings"

	asm "github.com/hsiaosiyuan0/ripplet/internal/asm"
)

type Closure struct {
	fn     *asm.FnShape
	args   *Args
	uppers map[string]*Object
}

func NewClosure(fn *asm.FnShape) *Closure {
	return &Closure{
		fn:     fn,
		uppers: make(map[string]*Object),
	}
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

var tureObj = &Object{
	Typ:  BOOL,
	Data: true,
}

var falseObj = &Object{
	Typ:  BOOL,
	Data: false,
}

func (a *Args) Get(i int) *Object {
	if i < a.Len {
		return a.slots[i]
	}
	return nilObj
}

func (a *Args) shift() *Object {
	if len(a.slots) == 0 {
		return nilObj
	}

	arg, rest := a.slots[0], a.slots[1:]
	a.slots = rest
	return arg
}

type GoFnImpl func(*Args) interface{}

type GoFn struct {
	Name string
	Impl GoFnImpl
}

type CallFrame struct {
	*Closure

	locals []*Object
	prev   *CallFrame

	pc   int
	base int
}

func NewCallFrame(closure *Closure) *CallFrame {
	frame := &CallFrame{
		locals: make([]*Object, closure.fn.LocalCnt),
	}
	frame.Closure = closure
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

func (v *Vm) GetChunk() *asm.Chunk {
	return v.chunk
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

		local := v.cf.locals[i]
		opd := v.popOperand()
		if local == nil {
			v.cf.locals[i] = opd
		} else {
			// keep reference
			local.Typ = opd.Typ
			local.Data = opd.Data
		}

	case asm.LOAD:
		i := v.fetchInstr()
		v.pushOperand(v.cf.locals[i])

	case asm.LOAD_EXT:
		i := v.fetchInstr()
		nameConst := v.chunk.GetConst(i)
		name := nameConst.Val.(string)
		v.pushOperand(v.externals[name])

	case asm.LOAD_ARG:
		i := v.fetchInstr()
		v.cf.locals[i] = v.cf.args.shift()

	case asm.CALL:
		argsLen := v.fetchInstr() - 1
		args := &Args{
			Len:   argsLen,
			slots: make([]*Object, argsLen),
		}
		for i := 0; i < argsLen; i++ {
			args.slots[i] = v.popOperand()
		}

		target := v.popOperand()
		if target.Typ == GFN {
			gfn := target.Data.(GoFnImpl)
			gfn(args)
		} else if target.Typ == CLOSURE {
			closure := target.Data.(*Closure)
			closure.args = args
			cf := NewCallFrame(closure)
			cf.base = len(v.opds)
			cf.prev = v.cf
			v.cf = cf
		} else {
			panic(fmt.Errorf("obj is not callable %v", target))
		}

	case asm.RET:
		// shrink opds for releasing those useless operands
		opds := make([]*Object, v.cf.base)
		copy(opds, v.opds)

		retVal := nilObj
		opdsLen := len(v.opds)
		if opdsLen > 0 {
			retVal = v.opds[opdsLen-1]
		}

		v.opds = append(opds, retVal)
		v.cf = v.cf.prev

	case asm.CONCAT:
		subsLen := v.fetchInstr()
		subs := make([]string, subsLen)

		for i := 0; i < subsLen; i++ {
			opd := v.popOperand()
			// TODO: toString
			ss := opd.Data.(string)
			subs = append(subs, ss)
		}

		v.pushOperand(&Object{
			Typ:  STR,
			Data: strings.Join(subs, ""),
		})

	case asm.CLOSURE:
		i := v.fetchInstr()
		shape := v.cf.fn.Subs[i]

		closure := &Closure{
			fn:     shape,
			uppers: make(map[string]*Object),
		}

		for name, loc := range shape.Upvals {
			if loc.Typ == asm.UP_LOC_LOCAL {
				i := loc.At.(int)
				local := v.cf.locals[i]
				if local == nil {
					local = &Object{
						Typ: NIL,
					}
					v.cf.locals[i] = local
				}
				closure.uppers[name] = local
			} else {
				closure.uppers[name] = v.cf.uppers[name]
			}
		}

		v.pushOperand(&Object{
			Typ:  CLOSURE,
			Data: closure,
		})

	case asm.LOAD_UP:
		name := v.chunk.ConstStr(v.fetchInstr())
		v.pushOperand(v.cf.uppers[name])

	case asm.STORE_UP:
		name := v.chunk.ConstStr(v.fetchInstr())
		v.cf.uppers[name] = v.popOperand()

	case asm.ADD:
		op2 := v.popOperand()
		op1 := v.popOperand()
		// TODO: implicitly data converting
		res := op1.Data.(float64) + op2.Data.(float64)
		v.pushOperand(&Object{
			Typ:  NUM,
			Data: res,
		})

	case asm.SUB:
		op2 := v.popOperand()
		op1 := v.popOperand()
		// TODO: implicitly data converting
		res := op1.Data.(float64) - op2.Data.(float64)
		v.pushOperand(&Object{
			Typ:  NUM,
			Data: res,
		})

	case asm.MUL:
		op2 := v.popOperand()
		op1 := v.popOperand()
		// TODO: implicitly data converting
		res := op1.Data.(float64) * op2.Data.(float64)
		v.pushOperand(&Object{
			Typ:  NUM,
			Data: res,
		})

	case asm.DIV:
		op2 := v.popOperand()
		op1 := v.popOperand()
		// TODO: implicitly data converting
		res := op1.Data.(float64) / op2.Data.(float64)
		v.pushOperand(&Object{
			Typ:  NUM,
			Data: res,
		})

	case asm.ARR:
		cnt := v.fetchInstr()
		arr := make([]*Object, cnt)
		for i := cnt - 1; i >= 0; i-- {
			arr[i] = v.popOperand()
		}
		v.pushOperand(&Object{
			Typ:  ARR,
			Data: arr,
		})

	case asm.SUBSCRIPT:
		// TODO: type assert
		idx := int(v.popOperand().Data.(float64))
		arr := v.popOperand().Data.([]*Object)
		// TODO: boundary check
		item := arr[idx]
		v.pushOperand(item)

	case asm.TEST:
		opd := v.popOperand()
		v.pushOperand(opd.AsBool())

	case asm.JMP_F:
		cond := v.popOperand()
		ofst := v.fetchInstr()
		if !cond.Data.(bool) {
			v.cf.pc += ofst
		}

	case asm.JMP:
		ofst := v.fetchInstr()
		v.cf.pc += ofst

	case asm.GT, asm.GE, asm.LT, asm.LE:
		op2 := v.popOperand()
		op1 := v.popOperand()

		switch op {
		case asm.GT:
			v.pushOperand(op1.GT(op2))
		case asm.GE:
			v.pushOperand(op1.GE(op2))
		case asm.LT:
			v.pushOperand(op1.LT(op2))
		case asm.LE:
			v.pushOperand(op1.LE(op2))
		}

	case asm.AND, asm.OR:
		op2 := v.popOperand().AsBool().Data.(bool)
		op1 := v.popOperand().AsBool().Data.(bool)

		if op == asm.ADD && op1 && op2 {
			v.pushOperand(tureObj)
		} else if op == asm.OR && (op1 || op2) {
			v.pushOperand(tureObj)
		} else {
			v.pushOperand(falseObj)
		}

	case asm.IS, asm.IS_NOT:
		op2 := v.popOperand().Data
		op1 := v.popOperand().Data

		if op == asm.IS && op1 == op2 {
			v.pushOperand(tureObj)
		} else if op == asm.IS_NOT && (op1 != op2) {
			v.pushOperand(tureObj)
		} else {
			v.pushOperand(falseObj)
		}

	case asm.NEG:
		op := v.popOperand()
		op.Data = -op.Data.(float64)
		v.pushOperand(op)

	case asm.BOOL_T:
		v.pushOperand(tureObj)

	case asm.BOOL_F:
		v.pushOperand(falseObj)

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
	v.cf = NewCallFrame(NewClosure(&v.chunk.Fn))

	for v.hasInstr() {
		op := v.fetchOp()
		v.dispatch(op)
	}
}
