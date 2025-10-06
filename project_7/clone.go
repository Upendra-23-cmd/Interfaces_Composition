package main

import "fmt"

// custum interface to clone data 
type Cloneable interface{
	Clone(data string)
}

// struct to read the cloned data
type Reader struct{}

// we read the cloned data
func (r *Reader) Clone(data string){
	fmt.Println("  cloned data is being read : ", data)
}

// wriiting in the cloned data and declaring a struct for it
type Writer struct{
	Name string
}

func (w *Writer) Clone(data string){
	fmt.Printf("[%s] Cloned data : %s\n",w.Name,data)
}

// preforming the cloneing function
func Doclone(c Cloneable , data string){
	c.Clone(data)
}


func main() {

	// Giving them the value
	Reader := &Reader{}
	Writer1 := &Writer{Name: "writer 1"}
	Writer2 := &Writer{Name: "writer 2"}

	// data being cloned
	data := "hello custum interface"
	data1 := "hello custum interface from writer 1"
	data2 := "hello custum interface from writer 2"

	// calling clone function
	Doclone(Reader, data)
	Doclone(Writer1,data1)
	Doclone(Writer2,data2)
}