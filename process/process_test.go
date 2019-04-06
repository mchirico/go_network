package process

import (
	"fmt"
	"testing"
)

func FunctionThingTest(t string) func(f *Thing) error {
	return func(f *Thing) error {
		f.inputType = t
		return nil
	}
}

func TestNewThing(t *testing.T) {
	g, _ := NewThing()
	fmt.Println(g)
}
