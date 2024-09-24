package main

import "fmt"

func imprimir() {
	fmt.Println("texto desdre la funcion imprimir")
}
func hola_string(s string) string {
	return "Hola " + s
}
func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("A Y B son iguales")
	} else {
		fmt.Println("A Y B no son iguales ")
	}
}
func main() {
	// imprimir()
	// fmt.Print("texto desde la fincion main")
	// çfmt.Println(hola_string(" mairongo"))
	// fmt.Println(sumar(5, 3))
	comparar_bool()
}
