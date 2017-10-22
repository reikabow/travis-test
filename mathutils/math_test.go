package mathutils

import (
	"testing"
)

func TestPowNegative(t *testing.T) {
	_, err := Pow(10, -10)
	if err == nil {
		t.Fail()
	}
}

func TestPowZero(t *testing.T) {
	n, err := Pow(10, 0)
	if err != nil {
		t.Fail()
	}
	if n != 1 {
		t.Fail()
	}
}

func TestPowPositive(t *testing.T) {
	n, err := Pow(10, 3)
	if err != nil {
		t.Fail()
	}
	if n != 1000 {
		t.Fail()
	}
}
func TestNegativePowOdd(t *testing.T) {
	n, err := Pow(-1, 3)
	if err != nil {
		t.Fail()
	}
	if n != -1 {
		t.Fail()
	}
}
func TestNegativePowEven(t *testing.T) {
	n, err := Pow(-1, 12)
	if err != nil {
		t.Fail()
	}
	if n != 1 {
		t.Fail()
	}
}
func TestNegativePowZero(t *testing.T) {
	n, err := Pow(-1, 12)
	if err != nil {
		t.Fail()
	}
	if n != 1 {
		t.Fail()
	}
}
