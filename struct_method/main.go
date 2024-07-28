package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (r authenticationInfo) getBasicAuth() string {
	result := fmt.Sprintf("Authorization: Basic %s:%s", r.username, r.password)
	return result
}
