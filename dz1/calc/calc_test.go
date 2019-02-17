package main

import "testing"

// сюда писать тесты

func TestAdd(t *testing.T) {
	expected := 3
	result, _ := Calculate("1 2 + =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestDifference(t *testing.T) {
	expected := 90
	result, _ := Calculate("100 20 10 - - =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestMulti(t *testing.T) {
	expected := 125
	result, _ := Calculate("5 5 5 * * =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestComplexFirst(t *testing.T) {
	expected := 15
	result, _ := Calculate("1 2 3 4 + * + =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestComplexSecond(t *testing.T) {
	expected := 21
	result, _ := Calculate("1 2 + 3 4 + * =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestBad(t *testing.T) {
	_, err := Calculate("1 2 + 3 s 4 + * ")
	if err == nil {
		t.Errorf("results not match\nGot: nill\nExpected some error")
	}
}

func TestDivision(t *testing.T) {
	expected := 2
	result, _ := Calculate("2 100 400 / / =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}