package main

func main() {
	//nesting for loop
	for i := 1; i <= 10; i++ {
		println(" main for: ", i)
		for j := 1; j <= 3; j++ {
			println("\t innet for: ", j)
		}
	}
}
