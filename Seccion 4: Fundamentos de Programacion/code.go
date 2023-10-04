package main

import (
	"fmt"
	"runtime"
)

var x int
var y uint8


func main() {
	// numeric_types()
	// str_types()
	// constants()
	// iota_const()
	bit_shifting()
}

func numeric_types(){
	x = 42
	y = uint8(x)

	fmt.Println(x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func str_types(){
	s1 := "Hello world"
	s2 := `This is a 
	broken line`

	fmt.Println(s1)
	fmt.Printf("%T\n", s1)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)
	
	fmt.Println()

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)
}

func constants(){
	const a = 42
	const b = 42.32
	const c = "Sebas Go"
	// const c string = "Sebas Go"

	type nombre string

	var d nombre
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	d = c
	fmt.Println(d)
}

func iota_const(){
	const(
		a = iota
		b
		c
	)

	const(
		d = iota
		e
		f
	)

	const(
		a1 = iota
		b1
		c1
		d1
		e1
		f1
	)

	fmt.Printf("El valor de a es: %v\n",a)
	fmt.Printf("El valor de b es: %v\n",b)
	fmt.Printf("El valor de c es: %v\n",c)
	fmt.Printf("El valor de d es: %v\n",d)
	fmt.Printf("El valor de e es: %v\n",e)
	fmt.Printf("El valor de f es: %v\n",f)
	fmt.Printf("El valor de a1 es: %v\n",a1)
	fmt.Printf("El valor de b1 es: %v\n",b1)
	fmt.Printf("El valor de c1 es: %v\n",c1)
	fmt.Printf("El valor de d1 es: %v\n",d1)
	fmt.Printf("El valor de e1 es: %v\n",e1)
	fmt.Printf("El valor de f1 es: %v\n",f1)
}

func bit_shifting(){

	// sin bit shifting
	// kb := 1024
	// gb := 1024 * kb
	// tb := 1024 * gb
	// fmt.Println("Sin bit shifting")
	
	// con bit shifting
	const(
		_ = iota
		kb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
		tb = 1 << (iota * 10)
	)
	fmt.Println("Con bit shifting")


	fmt.Printf("%d\t\t\t%b\n",kb,kb)
	fmt.Printf("%d\t\t\t%b\n",gb,gb)
	fmt.Printf("%d\t\t%b\n",tb,tb)
	
}
