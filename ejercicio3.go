package main
import "fmt"
const Pi = 3.1416
func circulo(radio float64) (area float64, perimetro float64) {
    area = Pi * radio * radio
    perimetro = 2 * Pi * radio
    return area, perimetro
}


// // func ecoDeLaMontana(mensaje string, iteraciones uint) {
// //     if iteraciones > 1 {
// //         ecoDeLaMontana(mensaje, iteraciones-1)
// //     }
// //     fmt.Println(mensaje, iteraciones)
// // }


// func circulo(radio float64) (area func() float64, perimetro func() float64) {

//     area = func() float64 {
//         return 3.1416 * radio * radio
//     }

//     perimetro = func() float64 {
//         return 2 * 3.1416 * radio
//     }

//     return
// }



func main() {
    a, p := circulo(8)
    // fmt.Println("El area del circulo es: ", a)
    // fmt.Println("El perimetro del circulo es: ", p)
// 	// ecoDeLaMontana("yodelayheehoo", 5)
// 	area, perimetro := circulo(10)
//     fmt.Println("El area del circulo es", area())
//     fmt.Println("El perimetro del circulo es", perimetro())
 }


