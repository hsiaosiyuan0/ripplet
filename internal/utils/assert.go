package utils

import (
	"fmt"
	"strings"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	match := false

	if as, ok := a.(string); ok {
		if bs, ok := b.(string); ok {
			match = strings.Compare(as, bs) == 0
		}
	} else {
		match = a == b
	}

	if match {
		return
	}

	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
