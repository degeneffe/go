package main

import (
	"fmt"
)
type Gorra struct {
  marca string
  color string
  precio float32
  plana bool
}

func main() {

var gorra_negra = Gorra{
  marca:"Nike",
  color:"Negro",
  precio: 25.50,
  plana:false}

  fmt.Println(gorra_negra.marca)

	var numero1 float32 = 10
	var numero2 float32 = 6
	fmt.Println("La suma es:")
	fmt.Println(operacion(numero1,numero2, "+"))

	fmt.Println("La resta es:")
	fmt.Println(operacion(numero1,numero2, "-"))

	fmt.Println("La multiplicacion es:")
	fmt.Println(numero1 * numero2)

	fmt.Println("La division es:")
	fmt.Println(numero1 / numero2)

  holaMundo()
  fmt.Println(devolverTexto())
}
func holaMundo(){
  fmt.Println("Hola Mundo")
}

func operacion(n1 float32, n2 float32 ,op string) float32 {
  var resultado float32
  if(op == "+"){
    resultado = n1 + n2
  }
  if(op == "-"){
    resultado = n1 - n2
  }

  return resultado
}

func devolverTexto() (string, string) {
  return "texto por defecto" , "Segundo Texto"
}
