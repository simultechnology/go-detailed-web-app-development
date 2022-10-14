package main

import "fmt"

type Animal interface {
	Bark() string
}

type Dog struct{}

func (d *Dog) Bark() string { return "Bow" }

type BullDog struct {
	Dog
}

type ShibaInu struct {
	Dog
}

func (s *ShibaInu) Bark() string { return "One" }

func DogVoice(d *Animal) string {
	return "hoge"
}

type Person struct {
	Name string
	Age  int
}

type Japanese struct {
	Person
	MyNumber int
}

func Hello(p Person) {
	fmt.Println("Hello " + p.Name)
}

func main() {
	fmt.Println("ch6-3 start!")

	bd := &BullDog{}
	fmt.Println(bd.Bark())
	var si Animal
	si = &ShibaInu{}
	fmt.Println(si.Bark())

	fmt.Println(DogVoice(&si))

	person := Person{
		"Taka", 45,
	}
	Hello(person)
}
