package model

type Project struct {
	Name        string      `json:"name"`
	Link        string      `json:"link"`
	Description string      `json:"description"`
	Stacks      []TechStack `json:"stacks"`
	Submitted   bool
}

type TechStack struct {
	Name string `json:"name"`
}
