// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)

package main

import "fmt"

// описание полей структуры

type Human struct {
	name    string
	surname string
	age     int
}

// методы, реализуемые структурой Human

func (h Human) Introduce() {
	fmt.Printf("Hi! My name is %s %s, I'm %d years old\n", h.name, h.surname, h.age)
}

func (h Human) Speak(smth string) {
	fmt.Printf("%s says `%s`\n", h.name, smth)
}

// Action наследует все поля и методы структуры Human
type Action struct {
	Human
}

func main() {
	// создание экземпляра структуры Action
	act := Action{
		Human{
			name:    "Alice",
			surname: "Armstrong",
			age:     32},
	}
	// инстанс структуры Action реализует методы, отнаследованные от Human
	act.Introduce()
	act.Speak("I'm happy))")
}
