package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true


func main() {
	// practice_1()
	// practice_2()
	// practice_3()
	// practice_4()
	practice_5()
}

// practice #1
func practice_1() {
	x := 42
	y := "James Bond"
	z := true

	// 1a
	fmt.Println(x, y, z)

	// 1b
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func practice_2(){
	fmt.Println(x, y, z)
}

func practice_3(){
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}

func practice_4(){
	type numero int
	var x numero
	
	fmt.Printf("El tipo de x es: %T\n", x)

	x = 42
	fmt.Println(x)
}

func practice_5(){
	type hotdog int
	var x hotdog
	var y int

	fmt.Println(x)
	fmt.Printf("El valor de 'x' es %v y su tipo es: %T\n", x, x)
	x = 2
	fmt.Println(x)

	y = int(x)
	fmt.Printf("El valor de 'y' es %v y su tipo es %T.\n", y, y)
}

