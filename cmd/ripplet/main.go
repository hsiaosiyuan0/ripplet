package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/hsiaosiyuan0/ripplet/internal/vm"
)

func main() {

	file := flag.String("f", "", "path of the script to run")
	dump := flag.Bool("d", false, "dump opcodes")
	flag.Parse()

	if *file == "" {
		panic("missing script, use `-f` to specify one")
	}

	var code []byte
	var err error
	if code, err = ioutil.ReadFile(*file); err != nil {
		panic(err)
	}

	print := &vm.GoFn{
		Name: "print",
		Impl: func(args *vm.Args) interface{} {
			for i := 0; i < args.Len; i++ {
				arg := args.Get(i)
				fmt.Println(arg.String())
			}

			return nil
		},
	}

	vm := vm.NewVm(nil)
	vm.AddExtFn(print)

	vm.SetCode(string(code))

	if *dump {
		fmt.Println(vm.GetChunk().Dump())
		return
	}

	vm.Exec()
}
