package _1201

import (
	"testing"
)

func Test01_1(t *testing.T) {
	actual, err := Run01One(get01OneData())
	expected := 12 + 38 + 15 + 77

	if err != nil {
		t.Errorf("Error occured: %s", err.Error())
	}

	if actual != expected {
		t.Errorf("Expect %d, but got %d", expected, actual)
	}
}

func get01OneData() []string {
	return []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
}
