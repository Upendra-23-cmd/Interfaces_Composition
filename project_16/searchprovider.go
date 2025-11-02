package main

import "fmt"


type SearchProvider interface {
	Search()
}

type Google struct{}
type Bing struct{}

func (g *Google) Search(){
	fmt.Println("We are searching on google search engine")
}

func (b *Bing) Search(){
	fmt.Println("We are searching on Bing search engine")
}

func main() {
	search := []SearchProvider{&Google{},&Bing{}}

	for _, sear := range search {
		sear.Search()
	}
}