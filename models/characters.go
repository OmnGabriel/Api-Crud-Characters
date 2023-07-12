package models

type Character struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Backstory string `json:"backstory"`
}

var Characters []Character
