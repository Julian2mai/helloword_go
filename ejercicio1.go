package main

import (
	"fmt"
	"reflect"
	"strconv"
)

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

func arreglo() {
	var aprendices [5]string
	aprendices[0] = "Mairongo"
	aprendices[1] = "Eduardo"
	aprendices[2] = "camilo"
	aprendices[3] = "julian"
	aprendices[4] = "andres"
	fmt.Println(aprendices[0])
}

func tipo_datos() {
	var texto string = "mairongo"
	var entero int = 1000
	var decimal float64 = 0.00001
	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(entero))
	fmt.Println(reflect.TypeOf(decimal))
}

func convertirganado (){
	var palabra string = "true"
	boolean, _ := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	}

func convernovia () {
	var palabra_bool bool =true
	strbool := strconv.FormatBool((palabra_bool))
	fmt.Println(strbool, reflect.TypeOf(strbool))
}

func main() {
	// imprimir()
	// fmt.Print("texto desde la fincion main")
	// fmt.Println(hola_string(" mairongo"))
	// fmt.Println(sumar(5, 3))
	// comparar_bool()
	// arreglo()
	// tipo_datos()
    // convertirganado()
	convernovia()
}
