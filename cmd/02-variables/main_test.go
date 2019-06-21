package main

import "testing"

func TestAddAge(t *testing.T) {
	testFriend := friend{
		age: 13,
	}
	testIncrement := 2
	result := addAge(testFriend, testIncrement)
	expected := 15
	if result != expected {
		t.Errorf("Failed Add Age. Expected: %d, Received: %d", expected, result)
	}
}

func TestCelebrateBirthday(t *testing.T) {
	testFriend := friend{
		age: 13,
	}
	celebrateBirthday(testFriend)
	result := testFriend.age
	expected := 14
	if result != expected {
		t.Errorf("Failed Celebrate Birthday. Expected: %d, Received: %d", expected, result)
	}
}
