package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	s := `[{"First":"Ali","Last":"Mohammadi","Age":90},{"First":"Reaza","Last":"Pishro","Age":15},{"First":"sos","Last":"Mast","Age":85}]`

	bs := []byte(s)

	// person := []person{}
	var person []person

	err := json.Unmarshal(bs, &person)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range person {
		fmt.Println(v)
	}
}
