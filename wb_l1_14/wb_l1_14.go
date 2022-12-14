// определить тип переменной: int, string, bool, channel из переменной типа interface{}

package main

import (
	"fmt"
	"reflect"
)

// определение типа переменной через конструкцию switch (не работает для каналов)
func switchMethod(vars []interface{}) {
	for _, v := range vars {

		switch t := v.(type) {
		case int:
			fmt.Println("int", t)
		case string:
			fmt.Println("string", t)
		case bool:
			fmt.Println("bool", t)
		default:
			fmt.Println("unknown")
		}
	}
}

// определение типа переменной через Sprintf()
func sprintMethod(vars []interface{}) {
	for _, v := range vars {
		t := fmt.Sprintf("%T", v)
		fmt.Println(t, v)
	}
}

// определение типа данных с помощью reflect.TypeOf()
func reflectionMethod(vars []interface{}) {
	for _, v := range vars {
		t := reflect.TypeOf(v)
		fmt.Println(t, v)
	}
}

func main() {
	var vars []interface{}
	ch := make(chan int)
	vars = append(vars, 42, "My string", true, ch)

	fmt.Println("Switch method")
	switchMethod(vars)
	fmt.Println()
	fmt.Println("Sprintf method")
	sprintMethod(vars)
	fmt.Println()
	fmt.Println("Reflection method")
	reflectionMethod(vars)
}
