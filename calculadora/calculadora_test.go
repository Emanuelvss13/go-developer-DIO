package calculadora

import "testing"

func ShouldSumCorrectly(t *testing.T) {
	value := sum(1, 2, 3)
	expected := 6

	if value != expected {

		t.Error("ShouldSumCorrectly")

	}
}

func ShouldSubtractCorrectly(t *testing.T) {
	value := subtration(100, 20, 40)
	expected := 40

	if value != expected {

		t.Error("ShouldSubtractCorrectly")

	}
}

func ShouldDivideCorrectly(t *testing.T) {
	value := division(9, 3)
	expected := 3

	if value != expected {

		t.Error("ShouldDivideCorrectly")

	}
}

func ShouldMultipleCorrectly(t *testing.T) {
	value := multiplication(10, 10, 5)
	expected := 500

	if value != expected {

		t.Error("ShouldMultipleCorrectly")

	}
}
