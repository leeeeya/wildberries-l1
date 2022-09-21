// Объединить данные значения температурных колебаний в группы с шагом в 10 градусов

package main

import "fmt"

func main() {
	a := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float32)

	for _, i := range a {
		m[int(i)/10*10] = append(m[int(i)/10*10], i)
	}
	fmt.Println(m)
}
