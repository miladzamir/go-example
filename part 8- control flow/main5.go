package main

import "fmt"

func main() {

	//print even number
	count := 0
	for {
		if count > 100000000000000 {
			break
		}

		if count%2 != 0 {
			count++
			//use continue to ignore remaning code
			continue
		}

		fmt.Println(count)
		count++
	}
}
