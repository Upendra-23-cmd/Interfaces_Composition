package main

import "fmt"

type Command interface{
	Execute()
}

type StartCommand struct{}
type StopCommand struct{}
type StatusCommand struct{}

func (s *StartCommand) Execute(){
	fmt.Println("Starting the process ..........")
}

func (s *StatusCommand) Execute(){
	fmt.Println("Status : Running process")
}

func (s *StopCommand) Execute(){
	fmt.Println("Stoping the running process .......")
}

func main(){
	cmd := []Command{&StartCommand{},&StatusCommand{},&StopCommand{}}

	for _, cmd := range cmd {
		cmd.Execute()
	}
}