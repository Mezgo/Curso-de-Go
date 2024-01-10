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
	practice8()
	// practice9()
	// practice10()
}

func practice1() {
	/* Usando un COMPOSITE LITERAL:
	- Crea un ARREGLO el cual tenga 5 VALORES de TIPO int
	- Asigna VALORES a cada posición dada por los índices.
	- Itera con Range sobre el arreglo e imprime los valores.
	Usando el paquete fmt
	- Imprime el TIPO del arreglo
	*/

	x := []int{1, 2, 3, 4, 5}
	for i, el := range x {
		txt := fmt.Sprintf("%v is at index: %v", el, i)
		fmt.Println(txt)
	}
}

func practice2() {
	/* Usando un COMPOSITE LITERAL:
	- Crea un SLICE de TIPO int
	- Asigna 10 VALORES
	- Haz Range sobre el slice e imprime los valores.
	Usando format para imprimir
	- Imprime el TIPO del slice
	*/

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, el := range x {
		txt := fmt.Sprintf("%v is at: %v", el, i)
		fmt.Println(txt)
	}
	fmt.Printf("The type of the slice is: %T\n", x)
}

func practice3() {
	/* Usando el código del ejercicio anterior, usa SLICING para crear los
	siguientes nuevos slices el cual serán impresos:
	- [42 43 44 45 46]
	- [47 48 49 50 51]
	- [44 45 46 47 48]
	- [43 44 45 46 47]
	*/
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[0:5])
	fmt.Println(x[5:10])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

func practice4() {
	/* Sigue los siguientes pasos:
	Comienza con este slice
	- x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	Agrégale el siguiente valor
	- 52
	Imprime el slice
	En UNA DECLARACIÓN agrega al slice los siguientes valores
	- 53
	- 54
	- 55
	Imprime el slice
	Agregale al Slice los siguientes valores
	- y := []int{56, 57, 58, 59, 60}
	print out the slice
	*/

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func practice5() {
	/* Para BORRAR de un slice, usamos APPEND en conjunto con SLICING(dividir).
	Para este ejercicio sigue los siguientes pasos:
	Comienza con un slice
	- x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	Usa APPEND & SLICING para obtener los siguientes valores el cual se los
	debes asignar a una variable “y” y luego imprimir:
	- [42, 43, 44, 48, 49, 50, 51]
	*/

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func practice6() {
	/* Crea un slice para almacenar los nombres de todos los estados en los
	Estados Unidos de América. ¿Cuál es la longitud del slice? ¿Cuál es la
	capacidad del Slice? Imprime todos los valores con su  índice de posición,
	sin utilizar la cláusula range. Aquí la lista de los estados:
	*/

	states := []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`,
		`Colorado`, `Connecticut`, `Delaware`, `Florida`,
		`Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`,
		`Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`,
		`Maryland`, `Massachusetts`, `Michigan`, `Minnesota`,
		`Mississippi`, `Missouri`, `Montana`, `Nebraska`,
		`Nevada`, `NewHampshire`, `NewJersey`, `NewMexico`,
		`NewYork`, `NorthCarolina`, `NorthDakota`, `Ohio`,
		`Oklahoma`, `Oregon`, `Pennsylvania`, `RhodeIsland`,
		`SouthCarolina`, `SouthDakota`, `Tennessee`, `Texas`,
		`Utah`, `Vermont`, `Virginia`, `Washington`,
		`WestVirginia`, `Wisconsin`, `Wyoming`}

	fmt.Printf("The length of the slice is: %v\n", len(states))
	fmt.Printf("The capacity of the slice is: %v\n", cap(states))

	for el := 0; el < len(states); el++ {
		fmt.Printf("State %v: is located at index %v\n", states[el], el)
	}
}

func practice7() {
	/* Crear un slice de slice de string ([][]string). Almacena los siguientes
	valores en un slice multi-dimensional:

	- "Batman", "Jefe", "Vestido de negro"
	- "Robin", "Ayudante", "Vestido de colores"

	Haz range sobre los registros, luego sobre los datos de cada registro.
	*/

	heros := [][]string{{"Batman", "Boss", "Black suit"},
		{"Robin", "Support", "Colored suit"}}

	for _, el := range heros {
		fmt.Println(el)
		for i2, el2 := range el {
			fmt.Printf("\t%v has %v in it\n", i2, el2)
		}
	}
}

func practice8() {

}
