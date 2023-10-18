package main

import (
	"fmt"
)

func main() {
	// practice1()
	// practice2()
	// practice3()
	// practice4()
	// practice5()
	practice6()
}

func practice1() {
	/*
	imprimir un número en decimal, binario y hex
	*/

	n := 25

	fmt.Printf("La representacion decimal de %v es: \t\t%v\n", n,n)
	fmt.Printf("La representacion binaria de %v es: \t\t%b\n", n, n)
	fmt.Printf("La representacion hexadecimal de %v es: \t%#x\n", n, n)
}

func practice2(){
	/* Usando operadores logicos, escribe expresiones y asigna sus valores a variables. 
	Finalmente, imprime dichas variables*/

	a := (1==1.1)
	b := (1!=1.1)
	c := (1<=1.1)
	d := (1>=1.1)
	e := (1<1.1)
	f := (1>1.1)
	
	fmt.Println(a,b,c,d,e,f)
}

func practice3(){
	/* Crea constantes CON TIPO y SIN TIPO. 
	Imprime el valor de las mismas. */

	const(
		a int = 1
		b = 2
	)

	fmt.Printf(`
	El valor y tipo de "a" es %v %T.\n
	El valor y tipo de "b" es %v %T.\n
	
	`, a, a, b,b)
}

func practice4(){
	/* Escribe un programa que:
		- Asignar un int a una variable
		- Imprímelo en decimal, binario, y hex
		- Has shifts de bits de ese int una posición
		  a la izquierda y asigna eso a una
		  variable
		- Imprime esa variable en decimal, binario y 
		  hex
 */

	var a int = 15
	fmt.Printf("La representacion decimal de %v es: \t\t%v\n", a,a)
	fmt.Printf("La representacion binaria de %v es: \t\t%b\n", a, a)
	fmt.Printf("La representacion hexadecimal de %v es: \t%#x\n", a, a)

	a <<= 1
	fmt.Printf("La representacion decimal de %v es: \t\t%v\n", a,a)
	fmt.Printf("La representacion binaria de %v es: \t\t%b\n", a, a)
	fmt.Printf("La representacion hexadecimal de %v es: \t%#x\n", a, a)
}

func practice5(){
	/* 
		Crea una variable de tipo string usando un string
		literal no interpretado (raw string literal). 
		Imprímelo.
 	*/

	var sr string = `esta es una linea
	y esta es otra`

	fmt.Println(sr)
}

func practice6(){
	/* 
		Usando iota, crea 4 constantes para los PRÓXIMOS 4 años. 
		Imprime los valores de las constantes.
	 */

	 const(
		_ = iota
		a = iota + 2024 
		b = iota + 2024 
		c = iota + 2024 
		d = iota + 2024 
	 )

	 fmt.Println(a,b,c,d)
}