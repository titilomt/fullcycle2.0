package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)

	if result != 5 {
		t.Errorf("The result %d must be %d", result, 5)
	}
}

func TestSub(t *testing.T) {
	result := Sub(2, 3)

	if result != -1 {
		t.Errorf("The result %d must be %d", result, -1)
	}
}

func TestMult(t *testing.T) {
	result := Mult(2, 3)

	if result != 6 {
		t.Errorf("The result %d must be %d", result, 6)
	}
}

func TestDiv(t *testing.T) {
	result := Div(2, 2)

	if result != 1 {
		t.Errorf("The result %d must be %d", result, 1)
	}
}

func TestPower(t *testing.T) {
	result := Power(2, 3)

	if result != 8 {
		t.Errorf("The result %f must be %d", result, 8)
	}
}

func TestSayHello(t *testing.T) {
	result := SayHello()

	if result != "Hello" {
		t.Errorf("The result %s must be %s", result, "Hello")
	}
}
