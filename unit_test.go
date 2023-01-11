package main

import "testing"

func TestcommonString(t *testing.T) {

	input1 := "ahelloworld"
	input2 := "hello"
	actualString := commonString(input1, input2)

	expectedString := "lll"
	if actualString != expectedString {
		t.Errorf("Expected String %s is not same as"+"actual string %s .", expectedString, actualString)
	}
}
