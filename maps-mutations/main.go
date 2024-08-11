package main

import "fmt"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	if !ok {
		err = fmt.Errorf("not found")
		return false, err
	}
	if existingUser.scheduledForDeletion {

		delete(users, name)
		return true, nil
	}
	return false, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
