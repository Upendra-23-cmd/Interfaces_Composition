package main

import "fmt"



type Notifier interface{
	notify(message string)
}

type Email struct{}
type Sms struct{}
type Stack struct{}

func (e *Email) notify(message string){
	fmt.Println("message sent through email : ",message)
}
func (e *Sms) notify(message string){
	fmt.Println("message sent through Sms : ",message)
}
func (e *Stack) notify(message string){
	fmt.Println("message sent through Stack : ",message)
}

func main()  {

	notify := []Notifier{&Email{},&Sms{},&Stack{}}
	for _, message := range notify{
		message.notify("hello from upendra ")
	}
}