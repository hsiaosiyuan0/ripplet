package main

import (
	"fmt"

	vm "github.com/hsiaosiyuan0/ripplet/internal/vm"
)

func main() {
	code := `
a := 1
repeat {
	a = a + 1
  if a > 10 then break
}
print(a)
	`

	print := &vm.GoFn{
		Name: "print",
		Impl: func(args *vm.Args) interface{} {
			for i := 0; i < args.Len; i++ {
				arg := args.Get(i)
				fmt.Println(arg)
			}

			return nil
		},
	}

	vm := vm.NewVm(nil)
	vm.AddExtFn(print)

	vm.SetCode(code)
	fmt.Println(vm.GetChunk().Fn.Dump(vm.GetChunk()))

	vm.Exec()
}
