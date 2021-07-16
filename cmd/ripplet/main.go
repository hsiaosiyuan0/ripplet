package main

import (
	"fmt"

	vm "github.com/hsiaosiyuan0/ripplet/internal/vm"
)

func main() {
	code := `
	a := "hello world"
	b := 1
	print(a, b)
	`

	print := &vm.GoFn{
		Name: "print",
		Impl: func(args *vm.Args) interface{} {
			for i := 0; i < args.Len; i++ {
				fmt.Println(args.Get(i))
			}

			return nil
		},
	}

	vm := vm.NewVm(nil)
	vm.AddExtFn(print)

	vm.SetCode(code).Exec()
}
