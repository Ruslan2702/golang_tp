package main

import "testing"

// сюда писать тесты

func TestAdd(t *testing.T) {
	expected := 3
	result := Calculate("1 2 + =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestDifference(t *testing.T) {
	expected := 90
	result := Calculate("100 20 10 - - =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestMulti(t *testing.T) {
	expected := 125
	result := Calculate("5 5 5 * * =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestComplexFirst(t *testing.T) {
	expected := 15
	result := Calculate("1 2 3 4 + * + =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestComplexSecond(t *testing.T) {
	expected := 21
	result := Calculate("1 2 + 3 4 + * =")
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}