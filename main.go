package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Digite um dos números: ")
	fmt.Scan(&a)

	fmt.Print("Digite o outro número: ")
	fmt.Scan(&b)


	fmt.Println("A soma é: ", a + b)
	fmt.Println("A subtração é:", a - b)
	fmt.Println("A divisão é:", a / b)
	fmt.Println("A multiplicação é: ", a * b)
	fmt.Println("O resto da divisão é: ", a % b)

	a++
	fmt.Println("Incrementar a", a)

	if a > 0 && b > 0{
		fmt.Println("Números Positivos")
	}
}
