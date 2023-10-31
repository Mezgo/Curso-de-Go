package main

import "fmt"

func main() {

	// practice1()
	// practice2()
	// practice3()
	// practice4()
	// practice5()
	// practice6()
	// practice7()
	// practice8()
	// practice9()
	practice10()
}

func practice1() {
/*
Imprime todos los números del 1 al 10,000
*/

	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func practice2(){
	/* Imprime el rune code point de las letras del alfabeto en mayúsculas tres veces */
	total := 0
	for i:=65; i<=90; i++ {
		total += 1
		fmt.Println(i)
		for j:=1; j<=3; j++ {
			fmt.Printf("\t%#U\n",i)
		}
	}
	fmt.Println(total)
}

func practice3(){
	/* Crea un ciclo usando la siguiente sintaxis
		for condición { }
	Haz que imprima los años que has vivido. */
	age := 0
	for i := 1998; i <2023; i++ {
		age += 1
		fmt.Println(i)
	}
	fmt.Println("Mi edad es:", age, 2023-1998)
}

func practice4(){
	/* Crea un ciclo for usando esta sintaxis
		for { }
	Haz que imprima los años que has vivido.
	*/
	var i int = 1998
	var age int = 0
	
	for {
		if i <= 2023{
			fmt.Println(i)
		} else{
			break
		}
		i++
		age++
	}
	fmt.Println("Age:", age-1)
}

func practice5(){
	/* Imprime el resto o módulo, el cual es resultado de dividir entre 4 cada número en el rango de 10 y 100.
 */
	
	for i := 10; i < 101; i++ {
		fmt.Println(i%4)
	}
}

func practice6(){
	/* Crea un programa que muestre el “if statement” en acción. */

	for i:=0; i<=10;i++{
		if i%2==0{
			fmt.Println(i, "es par")
		} else {
			fmt.Println(i, "es impar")
		}
	}
}

func practice7(){
	/* Usando el ejercicio anterior, crea un programa que use “else if” y “else”. */

	for i:=1; i<=10;i++{
		if i%2==0{
			fmt.Println(i, "es multiplo de 2")
		} else if i%3 == 0{
			fmt.Println(i, "es multiplo de 3")
		} else{
			fmt.Println("No se de que multiplo es")
		}
	}
}

func practice8(){
	/* Crea un programa que use una declaración switch sin expresión especificada. */

	switch {
	case true:
		fmt.Println("Tienes toda la razon")
	case false:
		fmt.Println("Estas equivocado")
	default:
		fmt.Println("Ni idea")
	}
}

func practice9(){
	/* Crea un programa que use una declaración switch con la expresión de switch especificada como una variable de TIPO string y el IDENTIFICADOR “deporteFav”. */

	deporteFav := "Tenis"

	switch deporteFav{
	case "Tenis":
		fmt.Println(deporteFav, ", es mi deporte favorito")
	case "Futbol", "Ajedrez":
		fmt.Println(deporteFav, ", esta cerca pero no")
	default:
		fmt.Println(deporteFav, ", no estas ni cerca")
	}
}

func practice10(){
	/* Escribe el resultado de las siguientes declaraciones:
		fmt.Println(true && true) 
		fmt.Println(true && false) 
		fmt.Println(true || true) 
		fmt.Println(true || false) 
		fmt.Println(!true) */

		fmt.Println(true && true) 
		fmt.Println(true && false) 
		fmt.Println(true || true) 
		fmt.Println(true || false) 
		fmt.Println(!true)
}