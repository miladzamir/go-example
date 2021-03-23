package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//n, err := fmt.Println("Hello")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(n)

	
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("wassup !")
	io.Copy(f,r)


	//f, err := os.Open("a.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer f.Close()
	//
	//bs, err := ioutil.ReadAll(f)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(bs))
}
