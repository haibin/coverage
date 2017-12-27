package person_test

import (
	"testing"

	"github.com/haibin/coverage/person"
)

func TestAge(t *testing.T) {
	if person.Age() != 10 {
		t.Error("wrong age")
	}
}

func TestCity(t *testing.T) {
	if person.City() != "singapore" {
		t.Error("wrong city")
	}
}