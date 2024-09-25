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
	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Print("este es el error",err)
	}

func convernovia () {
	var palabra_bool bool =true
	strbool := strconv.FormatBool((palabra_bool))
	fmt.Println(strbool, reflect.TypeOf(strbool))
}

func variables() {
	var nombre1, apellido1 string = "Carmen", "Electra"
	fmt.Println(nombre1, apellido1)

	var (
		nombre2    string = "Paul"
		edad2      int    = 32
		pensionado bool   = false
	)
	fmt.Println("Nombre: ", nombre2)
	fmt.Println("Edad: ", edad2)
	fmt.Println("Pensionado: ", pensionado)
}

func valor_cero() {
	var nombre string
	var edad int
	var peso float64
	var estudiante bool
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Edad: ", edad)
	fmt.Println("Peso: ", peso)
	fmt.Println("Estudiante: ", estudiante)
}

func declaracion_corta_variables() {
	nombre := "Benjamin Button"
	edad := 99
	peso := 80
	estudiante := false
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Edad: ", edad)
	fmt.Println("Peso: ", peso)
	fmt.Println("Estudiante: ", estudiante)
}

var var1 = "Este es el nivel 1"

func Scope() {
	var var2 = "Este es el nivel 2"
	{
		var var3 = "Este es el nivel 3"
		fmt.Println(var3)
	}
	fmt.Println(var1)
	fmt.Println(var2)
}

func Uso_punteros() {
	vehiculo1 := "rojo"
	fmt.Println("El vehiculo1 es", vehiculo1)

	vehiculo2 := vehiculo1
	fmt.Println("El vehiculo2 es", vehiculo2)

	vehiculo3 := &vehiculo1
	fmt.Println("El vehiculo3 es", *vehiculo3)

	vehiculo1 = "gris"

	fmt.Println("El vehiculo1 es", vehiculo1, &vehiculo1)
	fmt.Println("El vehiculo2 es", vehiculo2, &vehiculo2)
	fmt.Println("El vehiculo3 es", *vehiculo3, vehiculo3)
}

func Valor_vs_Referencia(altura float32) float32 {
	altura = altura * 3.28
	return altura
}

var altura float32 = 1.70

const Pi = 3.1416

func area(radio float64) float64 {
	return Pi * radio * radio
}

func main() {
	// imprimir()
	// fmt.Print("texto desde la fincion main")
	// fmt.Println(hola_string(" mairongo"))
	// fmt.Println(sumar(5, 3))
	// comparar_bool()
	// arreglo()
	// tipo_datos()
    //convertirganado()
	// convernovia()
	// valor_cero()
	// declaracion_corta_variables()
	// Scope()
	// Uso_punteros()
	// fmt.Println("La altura es:", altura, "mts")
	// fmt.Println("Al envejecer:", Valor_vs_Referencia(altura), "mts")
	// fmt.Println("Despues de envejecer:", altura, "mts")
	// fmt.Println("El area de un circulo cuyo radio es 3 es: ", area(3))
	
}
