package main

import "fmt"

// interface containg three methods
type Database interface{
	Connect()
	Query(query string)
	Close()
}

// MySql connection

type MySql struct{
	Name string
}

func (m *MySql) Connect(){
	fmt.Println("Connecting to the database MYSQL .....")
	fmt.Println("Connected to SQL")
}

func (m *MySql) Query(query string){
	fmt.Println("Running the requested query")
}

func (c *MySql) Close(){
	fmt.Println("Connection is being closed ")
}


// Postgress connection 

type Postgres struct{
	Name string
}

func (p *Postgres) Connect(){
	fmt.Println("Connecting to the database POSTGRES SQL ......")
	fmt.Println("Connected to Postgress")
}

func (p *Postgres) Query(query string){
	fmt.Println("Running the requested query")
}

func (p *Postgres) Close(){
	fmt.Println("Connection is being closed")
}

// using the interface to call the three methods
func UseDb(db Database){
	db.Connect()
	db.Query("SELECT *  FROM user")
	db.Close()
}

// Calling main function
func  main(){

	// create proxy database
	mysql := &MySql{Name: "Userdb"}
	postgers := &Postgres{Name: "Postdb"}

	// using the data base
	fmt.Println("")
	UseDb(mysql)

	fmt.Println("")
	UseDb(postgers)

}
