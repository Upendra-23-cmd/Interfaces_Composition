package main

import "fmt"

// we declared an interface
type Animal interface{
	speak()string
}

// varoios struct define 
type Dog struct{}
type Cat struct{}
type Cow struct{}



// define a fuction name speck that use receiver struct to return a string 
func (c Cat) speak()string{
	return "meow"
}
func (c Cow) speak()string{
	return "ammmyyy"
}

func (d Dog) speak()string{
	return "wooh"
}



func main(){

	// define an interface
	animal := []Animal{Dog{},Cat{},Cow{}}

	for _, animal := range animal{

		// here we call the speak interface
	fmt.Println(animal.speak())

	}
}