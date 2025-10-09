package main

import "fmt"


type Sortable interface{
	sort(data []int)[]int
}

type Ascending struct{}

func (a *Ascending) sort(data []int)[]int{
   n := len(data)
   sort := make([]int,n)
   copy(sort , data)

   for i := 0;i<n;i++{
	for j:= i+1;j<n;j++{
		if sort[i] > sort[j]{
			sort[i],sort[j]=sort[j], sort[i]
		}
	}
   }
   return sort
}

type Decending struct{}

func (d *Decending) sort(data []int) []int{
	n:= len(data)
	sort := make([]int,n)
	copy(sort, data)

	for i := 0;i<n;i++{
		for j:= i+1;j<n;j++{
			if sort[i] < sort[j]{
				sort[i],sort[j]=sort[j], sort[i]
			}
		}
	   }
	   return sort
}

func main(){
	data := []int{2,45,3423,23,67,89,7985,34}

	var sort Sortable

	sort = &Ascending{}
	fmt.Println("The Ascending order is : ",sort.sort(data))

	sort = &Decending{}
	fmt.Println("The Decending order is : ", sort.sort(data))
}