package friend

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testFriend *Friend

func TestFriend(t *testing.T) {
	t.Run("testNewFriend", testNewFriend)
	t.Run("testGreetFriend", testGreetFriend)
	t.Run("testCelebrateBirthday", testCelebrateBirthday)
	t.Run("testLiesAboutAge", testLiesAboutAge)
	t.Run("testIncrementErrors", testIncrementErrors)
	t.Run("testBirthdayError", testBirthdayError)
}

func testNewFriend(t *testing.T) {
	expectedName := "Dad"
	expectedAge := 65
	testFriend = NewFriend(expectedName, expectedAge)
	assert.Equal(t, expectedAge, testFriend.GetAge())
	assert.Equal(t, expectedName, testFriend.GetName())

}

func testGreetFriend(t *testing.T) {
	assert.NotNil(t, testFriend)
	expectedGreeting := "Hello Dad, who is 65, years old.\n"
	result := testFriend.GreetFriend()
	assert.Equal(t, expectedGreeting, result)
}

func testCelebrateBirthday(t *testing.T) {
	assert.NotNil(t, testFriend)
	currentAge := testFriend.GetAge()
	expectedAge := currentAge + 1

	testFriend.CelebrateBirthday()
	result := testFriend.GetAge()
	assert.Equal(t, expectedAge, result)
}

func testLiesAboutAge(t *testing.T) {
	assert.NotNil(t, testFriend)
	actualAge := testFriend.GetAge()
	lieIncrement := -5
	testFriend.SetLiesAboutAge(lieIncrement)
	expectedAge := actualAge + lieIncrement
	assert.Equal(t, expectedAge, testFriend.GetAge())
}

func testIncrementErrors(t *testing.T) {
	// Test age < 0
	resultLower, errLower := addAge(5, -6)
	assert.Equal(t, 0, resultLower)
	assert.Error(t, errLower)

	resultUpper, errUpper := addAge(100, 11)
	assert.Equal(t, 0, resultUpper)
	assert.Error(t, errUpper)
}

func testBirthdayError(t *testing.T) {
	// Test age < 0
	reallyOldFriend := NewFriend("old", 110)
	assert.Panics(t, reallyOldFriend.CelebrateBirthday)
}