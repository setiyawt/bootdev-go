package main

func isValidPassword(password string) bool {
	length := len(password)
	if length < 5 || length > 12 {
		return false
	}

	hasLower := false
	hasUpper := false
	hasDigit := false

	for i := range password {
		if password[i] >= 'a' && password[i] <= 'z' {
			hasLower = true
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			hasUpper = true
		} else if password[i] >= '0' && password[i] <= '9' {
			hasDigit = true
		}
	}

	return hasLower && hasUpper && hasDigit
}
