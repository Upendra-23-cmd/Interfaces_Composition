package main


type Mediaplayer interface{
	audio()
	video()
}

type User struct{}

func (u *User) audio(){
	println("The user want to listen audio")
}

func (u *User) video(){
	println("The user watch video")
}

func main(){
	user := &User{}

	user.audio()
	user.video()
	
}