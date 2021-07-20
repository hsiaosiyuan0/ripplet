package vm

import (
	"fmt"
	"strings"
)

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

func (ot ObjType) String() string {
	switch ot {
	case STR:
		return "STR"
	case NUM:
		return "NUM"
	case BOOL:
		return "BOOL"
	case NIL:
		return "NIL"
	case DIC:
		return "DIC"
	case ARR:
		return "ARR"
	case CLOSURE:
		return "CLOSURE"
	case GFN:
		return "GFN"
	default:
		panic(fmt.Errorf("unreachable: %d", ot))
	}
}

type Object struct {
	Typ  ObjType
	Data interface{}
}

func (o *Object) IsIntegral() bool {
	if o.Typ != NUM {
		return false
	}
	v := o.Data.(float64)
	return v == float64(int(v))
}

func (o *Object) String() string {
	switch o.Typ {
	case STR:
		return o.Data.(string)
	case NUM:
		if o.IsIntegral() {
			return fmt.Sprintf("%d", int(o.Data.(float64)))
		}
		return fmt.Sprintf("%f", o.Data)
	case BOOL:
		return fmt.Sprintf("%t", o.Data)
	case NIL:
		return "nil"
	case ARR:
		var b strings.Builder
		b.Write([]byte("["))
		arr := o.Data.([]*Object)
		arrLen := len(arr)
		for i, obj := range arr {
			fmtStr := "%s, "
			if i == arrLen-1 {
				fmtStr = "%s"
			}
			fmt.Fprintf(&b, fmtStr, obj.String())
		}
		b.Write([]byte("]"))
		return b.String()
	default:
		return fmt.Sprintf("%s#%v", o.Typ.String(), o.Data)
	}
}

func (o *Object) Bool() *Object {
	switch o.Typ {
	case NUM:
		n := o.Data.(float64)
		if n != 0 {
			return tureObj
		}

	case STR:
		s := o.Data.(string)
		if len(s) > 0 {
			return tureObj
		}

	case BOOL:
		b := o.Data.(bool)
		if b {
			return tureObj
		}

	case ARR, DIC, CLOSURE, GFN:
		return tureObj

	}

	return falseObj
}

func (o *Object) GT(rhs *Object) *Object {
	if o.Typ != rhs.Typ {
		return falseObj
	}
	switch o.Typ {
	case STR:
		op1 := o.Data.(string)
		op2 := rhs.Data.(string)
		if strings.Compare(op1, op2) == 1 {
			return tureObj
		}

	case NUM:
		op1 := o.Data.(float64)
		op2 := rhs.Data.(float64)
		if op1 > op2 {
			return tureObj
		}
	}

	return falseObj
}

func (o *Object) GE(rhs *Object) *Object {
	if o.Typ != rhs.Typ {
		return falseObj
	}
	switch o.Typ {
	case STR:
		op1 := o.Data.(string)
		op2 := rhs.Data.(string)
		cmp := strings.Compare(op1, op2)
		if cmp == 1 || cmp == 0 {
			return tureObj
		}

	case NUM:
		op1 := o.Data.(float64)
		op2 := rhs.Data.(float64)
		if op1 >= op2 {
			return tureObj
		}
	}

	return falseObj
}

func (o *Object) LT(rhs *Object) *Object {
	if o.Typ != rhs.Typ {
		return falseObj
	}
	switch o.Typ {
	case STR:
		op1 := o.Data.(string)
		op2 := rhs.Data.(string)
		cmp := strings.Compare(op1, op2)
		if cmp == -1 {
			return tureObj
		}

	case NUM:
		op1 := o.Data.(float64)
		op2 := rhs.Data.(float64)
		if op1 < op2 {
			return tureObj
		}
	}

	return falseObj
}

func (o *Object) LE(rhs *Object) *Object {
	if o.Typ != rhs.Typ {
		return falseObj
	}
	switch o.Typ {
	case STR:
		op1 := o.Data.(string)
		op2 := rhs.Data.(string)
		cmp := strings.Compare(op1, op2)
		if cmp == -1 || cmp == 0 {
			return tureObj
		}

	case NUM:
		op1 := o.Data.(float64)
		op2 := rhs.Data.(float64)
		if op1 <= op2 {
			return tureObj
		}
	}

	return falseObj
}

func (o *Object) EQ(rhs *Object) *Object {
	if o.Typ != rhs.Typ {
		return falseObj
	}
	switch o.Typ {
	case STR:
		op1 := o.Data.(string)
		op2 := rhs.Data.(string)
		cmp := strings.Compare(op1, op2)
		if cmp == 0 {
			return tureObj
		}

	case NUM:
		op1 := o.Data.(float64)
		op2 := rhs.Data.(float64)
		if op1 == op2 {
			return tureObj
		}

	case BOOL:
		op1 := o.Data.(bool)
		op2 := rhs.Data.(bool)
		if op1 == op2 {
			return tureObj
		}

	case NIL:
		return tureObj
	}

	return falseObj
}

func (o *Object) MustArray() *Array {
	if o.Typ != ARR {
		panic(fmt.Errorf("object is not array: %v", o.Typ))
	}
	return o.Data.(*Array)
}
