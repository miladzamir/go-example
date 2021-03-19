package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	st := `[{"First":"James","Age":32},{"First":"MoneyPenny","Age":27},{"First":"M","Age":54}]`

	var user []user

	err := json.Unmarshal([]byte(st), &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
