package main

import "fmt"

type Character interface{
	Roles()
}

type Hero struct{}
type Villian struct{}
type SupportCharacter struct{}
type BossMonster struct{}

func (h *Hero) Roles(){
	fmt.Println("[Hero] I am hero of the game")
}

func (v *Villian) Roles(){
	fmt.Println("[Villian] I am Villian of the role")
}

func (s *SupportCharacter) Roles(){
	fmt.Println("[Support roles] We help the hero in completing the task")
}

func (b *BossMonster) Roles(){
	fmt.Println("[Boss monster] we a boss monster we stop hero after every 5 levels")
}

func main(){
	roles := []Character{&Hero{},&Villian{},&SupportCharacter{},&BossMonster{}}

	for _,role := range roles{
		role.Roles()
	}
}