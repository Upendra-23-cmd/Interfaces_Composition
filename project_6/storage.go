package main

import "fmt"

type Storage interface {
	Save(data string)
	Load() string
}

type File struct{}

func (f File) Save(data string) { fmt.Println("Data stored on file system : ", data) }
func (f File) Load() string     { return "data loaded from file.txt" }

type S3 struct{}

func (s S3) Save(data string) { fmt.Println("Data stored on S3 Bucket  : ", data) }
func (s S3) Load() string     { return "data loaded from S3 bucket" }

type Database struct{}

func (d Database) Save(data string) { fmt.Println("Data stored in data base : ", data) }
func (d Database) Load() string     { return "data loaded from file.txt" }

func PreformTheAction(s Storage, data string) {
	s.Save(data)
	fmt.Print(s.Load())
	fmt.Println()
}

func main() {
	file := File{}
	S3 := S3{}
	Db := Database{}

	PreformTheAction(file, "File Content")

	PreformTheAction(S3, "S3 storage data")

	PreformTheAction(Db, "Database data  Content being fetched")
}
