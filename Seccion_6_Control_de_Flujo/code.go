package main

import "fmt"

func main(){

	// ciclo_for()
	// ciclo_anidado()
	// variacion_for()
	// break_continue()
	// imprimiendo_ascii()
	// declaracion_if()
	switch_declaration()
}

func ciclo_for() {

	for i:=0;i<10;i++ {
		fmt.Println(i)
	}
}

func ciclo_anidado(){
	for i:=0;i<10;i+=2 {
		fmt.Printf("ciclo externo: %d\n",i)
		for j:=0;j<=10;j+=5 {
			fmt.Printf("\tciclo interno: %d\n",j)
		}
	}
}

func variacion_for(){
	i := 0
	for i<5{
		fmt.Println(i)
		i++
	}
	fmt.Println("/////////////////////////////")
	for {
		if i > 9 { 
			break 
		} else{
			fmt.Println(i)
			i++
		}
	}
}

func break_continue(){
	x:=1
	
	for {
		
		x++
		if x>100{
			break
		}
		if x%2 != 0{
			// x++
			continue // devuelve al inicio del ciclo
		}
		fmt.Println(x)
		
	}
}

func imprimiendo_ascii(){
	for i:=33; i <= 122; i++{
		fmt.Printf("%d\t%#x\t%#U\n",i,i,i)
	}
}

func declaracion_if(){
	if true{
		fmt.Println(1)
	}
	if false{
		fmt.Println(2)
	}
	if !true{
		fmt.Println(3)
	}
	if !false{
		fmt.Println(4)
	}
	
	x:= 24

	fmt.Printf("x: %d\n", x)

	if x==20{
		fmt.Println("El valor de x es 20")
	} else if x == 21{
		fmt.Println("El valor de x es 21")
	} else if x == 22{
		fmt.Println("El valor de x es 22")
	}else {
		fmt.Println("El valor de x no es 20, 21 o 22")
	}
}

func switch_declaration() {
	x:=1
	switch {
	case x==1:
		fmt.Println("Si imprime")
	case x==45:
		fmt.Println("No imprime")
	default:
		fmt.Println("default")
		
	}

	y:= "Manzana"

	switch y {
	case "Limon", "Pera":
		fmt.Println("Frutas Verdes")
	
	case "Manzana", "Cereza":
		fmt.Println("Frutas Rojas")
		fallthrough
	
	default:
		fmt.Println("No es fruta Verde ni Roja")
	}
}