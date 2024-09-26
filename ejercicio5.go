package main
import "fmt"

func sumar(numeros ...int) int {
    // el total inicial es 0
    total := 0
    // recorrer todos los numeros
    for _, numero := range numeros {
        // en cada iteración sumar al total el valor del numero
        total = numero + total
    }
    // retornar el valor total
    return total
}

func Else_if() {
    var juguete string
    fmt.Println("Elige persona, animal o cosa:")
    fmt.Scanln(&juguete)
    if juguete == "persona" {
        fmt.Println("El objeto es una persona")
    } else if juguete == "cosa" {
        fmt.Println("El objeto es una cosa")
    } else if juguete == "animal" {
        fmt.Println("El objeto es un animal")
    } else {
        fmt.Println("El objeto es otra categoria")
    }
}

func residuo () {
	var x int
    fmt.Println("Ingrese x:")
	fmt.Scanln(&x)
	if x%2 ==0{
		fmt.Println("El numero es par")
	}else {
		fmt.Println("El numero es impar")
	}
}

func main() {
    // fmt.Println(sumar( 8 ))
    // fmt.Println(sumar(5, 1))
    // fmt.Println(sumar(3,2, 5))
	// fmt.Println(sumar(3,2,5,6,7,8,8,8,8, 5))

	// Else_if()
	residuo()
}

