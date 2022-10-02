// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.

// 	для переменной v выделяется память и до завершения программы она не сможет освободиться,
//	так как на неё ссылается переменная justString.
//	в то же время используется только малая часть выделенной памяти, остальная занята бесполезно.
//	лучше выделять сразу столько памяти, сколько требуется

//	var justString string

//	func someFunc() {
//		v := createHugeString(1 << 10)
//		justString = v[:100]
//	}

//	func main() {
//		someFunc()
//	}

package main

import "fmt"

func createHugeString(len int) string {
	var str string
	for i := 0; i < len; i++ {
		str += "1"
	}
	return str
}

func someFunc() string {

	justString := createHugeString(100)
	return justString
}

func main() {
	fmt.Println(someFunc())
}
