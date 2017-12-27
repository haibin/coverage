package person

import (
	"testing"
)

func TestFoo(t *testing.T) {
	if foo() != "foo" {
		t.Error("wrong foo")
	}
}