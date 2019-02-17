package main

import (
	"fmt"
	"testing"
)

// сюда писать тесты

func TestBad(t *testing.T) {
	_, err := Calculate("1 2 + 3 s 4 + * ")
	if err == nil {
		t.Errorf("results not match\nGot: nill\nExpected some error")
	}
}

func TestCalc(t *testing.T) {
	testCases := []struct {
		expression  string
		expected  int
	}{
		{"1 2 + =", 3},
		{"100 20 10 - - =", 90},
		{"5 5 5 * * =", 125},
		{"1 2 3 4 + * + =", 15},
		{"1 2 + 3 4 + * =", 21},
		{"2 100 400 / / =", 2},
	}

	for idx, myCase := range testCases {
		t.Run(fmt.Sprintf("Test #%d: %s", idx, myCase.expression), func(t *testing.T) {
			result, _ := Calculate(myCase.expression)
			if result != myCase.expected {
				t.Errorf("results not match\nGot: %v\nExpected: %v", result, myCase.expected)
			}
		})
	}
}