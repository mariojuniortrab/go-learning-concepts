package math

import "testing"

const defaultError = "Espected value: %v, Value found: %v"

func TestAverage(t *testing.T) {
	espectedValue := 7.28
	value := average(7.2, 9.9, 6.1, 5.9)

	if value != espectedValue {
		t.Errorf(defaultError, espectedValue, value)
	}

}
