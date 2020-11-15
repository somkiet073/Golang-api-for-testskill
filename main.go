package main

import (
	"fmt"

	api "github.com/somkiet073/Golang-api-for-testskill/app"
	"golang.org/x/crypto/bcrypt"
)

func test() {
	password := []byte("MyDarkSecret")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err) // nil means it is a match

}

func main() {

	// https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121
	// test()
	api.Run()
}
