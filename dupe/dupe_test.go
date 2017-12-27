package dupe_test

import (
	"testing"

	"github.com/haibin/coverage/person"
)

func TestAge(t *testing.T) {
	if person.Age() != 10 {
		t.Error("wrong age")
	}
}
