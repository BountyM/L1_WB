package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

func (h *Human) Talk() {
	fmt.Println("Hi my name is " + h.Name + " " + h.Surname)
}

type Action struct {
	Human
	Movement string
}

func main() {
	h := Human{
		Name:    "Maxim",
		Surname: "Kuzin",
	}
	h.Talk()

	a := Action{
		Human:    Human{"Max", "Kuzin"},
		Movement: "goes",
	}
	// Action имеет все методы Human
	a.Talk()

}
