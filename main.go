package main

import (
	"fmt"

	"github.com/somkiet073/Golang-api-for-testskill/app/auth"
)

func main() {
	// https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121

	token, err := auth.CreateToken(21)

	fmt.Println(token, err)
}
