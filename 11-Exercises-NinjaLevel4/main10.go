package main

import "fmt"

func main() {
	//delete record

	x := map[string][]string{
		"bound_james": []string{
			"shaken", "martini", "women",
		},
		"miss_moneypeny": []string{
			"notela", "coffe", "men",
		},
	}

	x["rez_ahmadi"] = []string{
		"daf", "asd", "21d",
	}

	if _, ok := x["miss_moneypeny"]; ok {
		delete(x, "miss_moneypeny")

	} else {
		fmt.Println("not exsist")
	}

	fmt.Println(x)

	for i, v := range x {
		fmt.Println("index : ", i)
		for _, d := range v {
			fmt.Println(d)
		}
	}
}
