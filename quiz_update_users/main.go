package main

import "fmt"

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

func newUser(name string, membershipType membershipType) User {
	var limitMessages int

	if membershipType == TypePremium {
		limitMessages = 1000
	} else {
		limitMessages = 100
	}

	return User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: limitMessages,
		},
	}
}

func main() {
	user := newUser("Renarin", TypeStandard)
	fmt.Println(user)
}
