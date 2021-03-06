package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func Shell(command string, args ...string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(command, args...)
	fmt.Print(cmd.String())

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		log.Fatalf("antlr lexer error %s", stderr.String())
	}
	fmt.Println(stdout.String())
}
