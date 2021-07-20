package vm

import (
	"fmt"
	"testing"
)

var assert = &GoFn{
	Name: "assert",
	Impl: func(args *Args) interface{} {
		op1 := args.Get(0)
		op2 := args.Get(1)
		eq := op1.EQ(op2).Data.(bool)
		if !eq {
			panic(fmt.Errorf("not eq %v %v", op1, op2))
		}
		return nil
	},
}

func runAssert(code string) {
	vm := NewVm(nil)
	vm.AddExtFn(assert)
	vm.SetCode(code)
	vm.Exec()
}

func TestCallNative(t *testing.T) {
	runAssert(`
  a := "hello world"
  assert(a, "hello world")
`)
}

func TestCallFn(t *testing.T) {
	runAssert(`
  f := (name) => {
    name
  }
  
  assert(f("hello world"), "hello world")
`)
}

func TestRet(t *testing.T) {
	runAssert(`
  f := (a) => {
    a
    2
  }
  
  a := f()
  assert(a, 2)
  `)
}

func TestUpvalDirect(t *testing.T) {
	runAssert(`
  b := 2
  f := (a) => {
    b
  }
  
  a := f()
  assert(a, 2)
  `)
}

func TestUpvalIndirect(t *testing.T) {
	runAssert(`
  b := 2
  f := (a) => {
    // implicit capture 'b'
    () => {
      b
    }
  }
  
  a := f()
  assert(a(), 2)
  `)
}

func TestUpvalNest(t *testing.T) {
	runAssert(`
  b := 2
  f := (a) => {
    b = 3
    () => {
      b
    }
  }
  
  a := f()
  assert(a(), 3)
  `)
}

func TestArray(t *testing.T) {
	runAssert(`
  a := [1, 2, 3]
  assert(a[0], 1)
  assert(a[1], 2)
  assert(a[2], 3)
  `)
}

func TestIf(t *testing.T) {
	runAssert(`
  a := 1
  b := if a then 2 else 3
  assert(b, 2)
  `)
}

func TestIfRet(t *testing.T) {
	runAssert(`
  a := 1
  f := () => {
    if a >= 1 then 2 else 3
  }
  assert(f(), 2)
  `)
}

func TestIfBlockRet(t *testing.T) {
	runAssert(`
  a := 1
  b := if a then {
    if a >= 1 then {
      5
    } else 6
  } else 3
  assert(b, 5)
  `)
}

func TestMath(t *testing.T) {
	runAssert(`
  a := 1
  b := 2
  assert(a + b, 3)
  assert(a - b, -1)
  assert(a * b, 2)
  assert(a / b, 0.5)
  assert(a + b * 2, 5)
  `)
}

func TestRepeat(t *testing.T) {
	runAssert(`
  a := 1
  repeat {
    a = a + 1
  } until a > 10
  assert(a, 11)
  `)
}

func TestRepeatBrk(t *testing.T) {
	runAssert(`
  a := 1
  repeat {
    a = a + 1
    break
  }
  assert(a, 2)
  `)
}

func TestRepeatBrkInIf(t *testing.T) {
	runAssert(`
  a := 1
  repeat {
    a = a + 1
    if a > 10 then break
  }
  assert(a, 11)
  `)
}

func TestRelation(t *testing.T) {
	runAssert(`
  a := 1
  b := 2
  assert(a > b, false)
  assert(a <= 1, true)
  `)
}

func TestEq(t *testing.T) {
	runAssert(`
  a := 1
  b := 1
  assert(a is b, true)

  c := 2
  assert(c isnt b, true)
  `)
}

func TestCalRec(t *testing.T) {
	runAssert(`
  f := (n) => {
    if n is 0 then 0
    else if n is 1 then 1
    else f(n - 1) + f(n - 2)
  }
  assert(f(10), 55)
  `)
}

func TestStringInterp(t *testing.T) {
	runAssert(`
  a := "hello world"
  b := 1
  assert("a: {a} {b}", "a: hello world 1")
  `)
}
