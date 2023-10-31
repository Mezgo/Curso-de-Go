package main

import "fmt"

func main() {
	// arrays()
	// slice()
	// multidimensional_slice()
	maps()
}

func arrays() {
	var arr [5]int
	arr[2] = 13
	a := fmt.Sprintf("arr es un arreglo de tipo %T.\nTiene largo %v\nSu contenido es %v", arr, len(arr), arr)

	fmt.Println(a)
}

func slice() {
	// Composite Literal
	x := []int{2, 3, 4, 5}
	out := fmt.Sprintf("%T", x)
	fmt.Println(x)
	fmt.Println(out)

	// Iterate on slice
	// 	for range
	fmt.Println(cap(x))
	for i, el := range x {
		txt := fmt.Sprintf("%v esta ubicado en el indice: %v", el, i)
		fmt.Println(txt)
	}

	// Slicing a slice
	fmt.Println(x[0:3])

	// add elements to slice
	x = append(x, 6, 7, 8, 9)
	fmt.Println(x)

	y := []int{10, 11, 12, 13}
	x = append(x, y...)
	fmt.Println(x)

	// delete elements from slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// make()
	new_x := make([]int, 10, 14)
	fmt.Println(new_x)
	fmt.Println(cap(new_x))
	new_x[0] = 14
	fmt.Println(new_x)
	fmt.Println(cap(new_x))

	new_x = append(x[0:], new_x...)
	fmt.Println(new_x)
	fmt.Println(cap(new_x))
}

func multidimensional_slice() {
	et := []string{"Sebastian", "Gomez", "Tenis", "Futball", "F1"}
	fmt.Println(et)

	jl := []string{"Karen", "Bolanos", "Ultimate", "Basketball", "Volleyball"}
	fmt.Println(jl)

	el := [][]string{et, jl}
	fmt.Println(el)
}

func maps() {
	// Creacion
	m := map[string]int{
		"Seb":   24,
		"Karen": 25,
	}
	fmt.Println(m)
	
	// Access values through key 
	fmt.Println(m["Karen"])

	// Try to access a non existent key
	fmt.Println(m["Daniela"])

	// Validate if a key exists
	v, ok := m["Daniela"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Daniela"]; !ok{
		fmt.Println("El elemnto no existe:", v)
	}

	// Add elements to the map
	m["Daniela"] = 20
	fmt.Println(m)

	// Loop through the map
	for k, v := range m{
		fmt.Printf("%v is %v years old\n", k, v)
	}

	// Delete elements from the map
	delete(m, "Karen") // NOTE: This does not check if the key exists in the map
	fmt.Println(m)

	m["Karen"] = 25
	// To solve the previous problem
	if v, ok := m["Karen"]; ok{
		txt := fmt.Sprintf("La llave con valor %v fue eliminada",v)
		delete(m, "Karen")
		fmt.Println(txt)
	}else{
		fmt.Println("The specified key doesn't exist")
	}
}
