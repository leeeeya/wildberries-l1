package main

import "fmt"

type Human struct {
	name    string
	surname string
	age     int
}

func (h Human) Introduce() {
	fmt.Printf("Hi! My name is %s %s, I'm %d years old\n", h.name, h.surname, h.age)
}

func (h Human) Speak(smth string) {
	fmt.Printf("%s says `%s`\n", h.name, smth)
}

type Action struct {
	Human
}

func main() {
	act := Action{
		Human{
			name:    "Alice",
			surname: "Armstrong",
			age:     32},
	}
	act.Introduce()
	act.Speak("I'm happy))")
}
