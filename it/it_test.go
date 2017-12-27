package it_test

import (
	"testing"

	"github.com/haibin/coverage/person"
)

func TestName(t *testing.T) {
	if person.Name() != "john" {
		t.Error("wrong name")
	}
}
