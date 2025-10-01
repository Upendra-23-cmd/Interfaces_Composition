package main

import (
	"fmt"
	"os"
)

// Logger interface
type Logger interface{
	Log(message string)
}

type ConsoleLogger struct {}

func (c *ConsoleLogger) Log(message string){
	fmt.Println("[console]",message)
}

type FileLogger struct{
	File *os.File
}

func NewFileLogger(filename string) *FileLogger{
	file, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY , 0644 )
	if err != nil {
		panic(err)
	}
	return &FileLogger{File:file}
}

func (f *FileLogger) Log(message string){
	f.File.WriteString(message + "\n")
}

func(f *FileLogger) Close(){
	f.File.Close()
}

func main() {

	console := ConsoleLogger{}
	console.Log("Hello, console !")

	file := NewFileLogger("log.txt")
	defer file.Close()
	file.Log("hello file")

	fmt.Println("Check for the file log entry")
}
