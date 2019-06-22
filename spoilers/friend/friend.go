package friend

import (
	"errors"
	"fmt"
)

// Friend holds the properties for tracking my relationships.
type Friend struct {
	name         string
	age          int
	ageOffset    int
	liesAboutAge bool
}

// NewFriend initializes a returns a new friend
func NewFriend(name string, age int) *Friend {
	return &Friend{
		name: name,
		age:  age,
	}
}

// SetLiesAboutAge helps me remeber if a person wants to be acknowledged with a different age than they actually are.
func (f *Friend) SetLiesAboutAge(offset int) {
	f.liesAboutAge = true
	f.ageOffset = offset
}

// GreetFriend prints a greeting for name and age of friend
func (f *Friend) GreetFriend() string {
	return fmt.Sprintf("Hello %s, who is %d, years old.\n", f.name, f.age)
}

// CelebrateBirthday increments the friends age by one
func (f *Friend) CelebrateBirthday() {
	newAge, err := addAge(f.age, 1)
	if err != nil {
		panic(err)
	}
	f.age = newAge
}

// GetAge gets friends "age" with any offset they lie about
func (f *Friend) GetAge() int {
	if !f.liesAboutAge {
		return f.age
	}

	fakeAge, _ := addAge(f.age, f.ageOffset)
	return fakeAge
}

// GetName gets the friend's name
func (f *Friend) GetName() string {
	return f.name
}

// addAge checks current age and increment and returns new age, will return error if new age is too big or too small.
func addAge(currentAge int, increment int) (int, error) {

	newAge := currentAge + increment

	if newAge > 110 || newAge < 0 {
		return 0, errors.New("invalid age")
	}
	return newAge, nil
}
