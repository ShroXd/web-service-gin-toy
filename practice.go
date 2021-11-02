package main

import "fmt"

type person struct {
	Name string
	Age  int
}

var persons = []person{
	{
		Name: "Spike",
		Age:  26,
	},
	{
		Name: "Jet",
		Age:  29,
	},
}

var spike = person{
	Name: "Spike",
	Age:  26,
}

func main() {
	//sayHello(persons)
	switch t := interface{}(spike).(type) {
	default:
		fmt.Println("nice")
	case person:
		fmt.Println("He is a person!", t)

	}
}

func sayHello(arr []person) {
Nice:
	for _, p := range arr {
		switch p.Name {
		case "Jet":
			break Nice
		case "Spike":
			fmt.Println("Nice!!")
		}
	}
}
