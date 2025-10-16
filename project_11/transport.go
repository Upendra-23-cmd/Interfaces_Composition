package main

import "fmt"

type Transport interface {
	mode()
}

type Car struct{}
type Bike struct{}
type Plane struct{}

func (c *Car) mode() {
	fmt.Println("The person is travelling by car")
}

func (b *Bike) mode() {
	fmt.Println("The person is travelling by Bike")
}

func (p *Plane) mode() {
	fmt.Println("The person is travelling by plane")
}

func main() {
	transport := []Transport{&Car{}, &Bike{}, &Plane{}}

	for _, count := range transport {
		count.mode()
	}
}
