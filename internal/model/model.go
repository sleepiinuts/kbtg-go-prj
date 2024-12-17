package model

type Request struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Address struct {
		No   int    `json:"no"`
		Road string `json:"road"`
	} `json:"address"`
}
