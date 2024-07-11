package model

type Project struct {
	Name        string   `json:"name"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	Stacks      []string `json:"stacks"`
	Submitted   bool     `json:"-"`
}
