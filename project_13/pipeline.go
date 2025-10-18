package main

import "fmt"

type Pipeline interface{
	stage()
}

type Build struct{}
type Test struct{}
type Deploy struct{}

func ( b *Build) stage(){
	fmt.Println("The project is building ...")
}

func ( t *Test) stage(){
	fmt.Println("The project is being tested ...")
}

func (d *Deploy) stage(){
	fmt.Println("The Project is being deployed ...")
}

func main(){

	stage := []Pipeline{&Build{},&Test{},&Deploy{}}

	for _, stage := range stage{
		stage.stage()
	}
}