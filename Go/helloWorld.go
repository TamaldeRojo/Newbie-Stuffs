package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	// Hola mundo

	fmt.Println("Hola bb")
	// Variables

	var myVar string //= "Go has a compiler lol"
	//fmt.Println(myVar)
	myVar = "can i re assigne the value?"
	fmt.Println(myVar)

	var myInt = 7
	fmt.Println(myInt)
	fmt.Println(reflect.TypeOf(myInt))

	var myFloat float32 = 1.1
	fmt.Println(myFloat + float32(myInt))

	myNewString := "soy un string"
	fmt.Println(myNewString)

	const myConst = "soy una constante lol"
	fmt.Println(myConst)

	if myInt != 10 {
		myInt = 10
		fmt.Println(" ya es 10 pa")
	}
	if myInt == 10 {
		fmt.Println("Hola pa")
	}
	if myInt == 10 && myFloat == 1.1 {
		fmt.Println("Lol")
	}
	if myInt == 10 || myNewString != "queso" {
		fmt.Println("Ponte a hacer algo pa")
	}

	//Data structures
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray)
	var myArray2 [3]string
	fmt.Println(myArray2)

	myMap := make(map[string]int)
	myMap["Josue"] = 21
	myMap["Carlos"] = 56
	myMap["Joel"] = 12
	fmt.Println(myMap)

	//List
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList)

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
	for key, value := range myMap {
		fmt.Println(key, value)
	}
	fmt.Println(myFunction())

	//no classes lol

	type MyStruct struct {
		name string
		age  int
	}
	myStruct := MyStruct{"Josue", 21}
	fmt.Println(myStruct)
}

func myFunction() string {
	fmt.Println(" Una funciooooonn loool")
	return "hola"
}
