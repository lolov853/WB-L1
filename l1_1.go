package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Salary int
}

func (h Human) ShortBio() {
	fmt.Println(h.Name)
	fmt.Println(h.Age)
	fmt.Println(h.Salary)
}

type Action struct {
	info       Human
	Occupation string
}

func main() {
	action := Action{
		info: Human{
			Name:   "Andrey",
			Age:    20,
			Salary: 70000,
		},
		Occupation: "data engineer",
	}
	action.info.ShortBio()

	fmt.Println(action.Occupation, "\n", action.info.Name, "\n", action.info.Age, "\n", action.info.Salary)

}
