package process

import (
	"fmt"
	"testing"
)

func TestNewThing(t *testing.T) {

	g, _ := NewThing()
	fmt.Println(g)

}
