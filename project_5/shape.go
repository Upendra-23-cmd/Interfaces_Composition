package main

import "fmt"

type Shape interface{
	shape()
}

type Rectangle struct{}
type Circle struct{}
type Triangle struct{}

func (R Rectangle) shape(){
	fmt.Println("this is rectangle")
}

func (C Circle) shape(){
	fmt.Println("This is circle")
}

func (T Triangle) shape(){
	fmt.Println("This is triangle")
}

func main(){
	shape := []Shape{Rectangle{},Circle{},Triangle{}}
	for _, shape := range shape {
		shape.shape()
	}
}
