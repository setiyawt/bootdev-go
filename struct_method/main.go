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

/*
package main

import (
	"fmt"
)

// Definisi struct Person
type Person struct {
	firstName, lastName string
	age                 int
}

// Metode untuk mengembalikan nama lengkap
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// Metode untuk mengembalikan apakah seseorang adalah dewasa
func (p Person) isAdult() bool {
	return p.age >= 18
}

// Metode untuk menambahkan tahun ke umur
func (p *Person) addYear() {
	p.age++
}

func main() {
	// Membuat instance dari Person
	p := Person{firstName: "John", lastName: "Doe", age: 20}

	// Memanggil metode fullName
	fmt.Println("Full Name:", p.fullName()) // Output: Full Name: John Doe

	// Memanggil metode isAdult
	fmt.Println("Is Adult:", p.isAdult()) // Output: Is Adult: true

	// Menambahkan tahun ke umur
	p.addYear()
	fmt.Println("New Age:", p.age) // Output: New Age: 21
}


*/
