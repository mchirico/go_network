package process

import (
	"fmt"
	"github.com/mchirico/go_network/yamlpkg"
	"log"
	"os"
)

type ReturnType []string
type FunctionConfig func(string) (ReturnType, error)

func A(a string) (ReturnType, error) {

	r := append(ReturnType{}, a)
	r = append(r, "done")

	return r, nil
}

type Thing struct {
	functionConfig FunctionConfig
	inputType      string
}

func NewThing(options ...func(*Thing) error) (ReturnType, error) {
	f := &Thing{}

	f.inputType = "HOME"
	f.functionConfig = func(b string) (ReturnType, error) {

		home := os.Getenv(b)

		file := home + "/.networkScriptConfig.yaml"
		fmt.Println(file)

		c := yamlpkg.Config{}
		err := c.Read(file)
		if err != nil {
			log.Fatalf("Not able to read .networkScriptConfig.yaml: %v\n", err)
			return []string{}, err
		}
		return c.ListGroups(), nil

	}

	for _, op := range options {
		err := op(f)
		if err != nil {
			return ReturnType{}, err
		}
	}
	return f.functionConfig(f.inputType)
}

func OptionalFn(f *Thing) error {
	f.functionConfig = A
	return nil
}

func OptionalFunctionThing(t string) func(f *Thing) error {
	return func(f *Thing) error {
		f.inputType = t
		return nil
	}
}
