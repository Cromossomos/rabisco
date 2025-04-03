package main

import "fmt"

func main() {
	var ages = [4]int{17, 20, 40, 87}
	nomes := [4]string{"Mario", "Deadpool", "Hermes", "Coiso"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[3] = "Algo"
	fmt.Println(nomes)
}
