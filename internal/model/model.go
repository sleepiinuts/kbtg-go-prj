package model

type Request struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Address struct {
		No   int    `json:"no"`
		Road string `json:"road"`
	} `json:"address"`
}

type School struct {
	Rooms []Room
}

type Room struct {
	No       int
	Students []Student
}

type Student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
