package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))

	pw := `password123`
	// pw := `password12333`

	err = bcrypt.CompareHashAndPassword(bs, []byte(pw))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("cool!")

}
