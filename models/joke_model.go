package models

//Joke is the model of joke fetch from server
type Joke struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}
