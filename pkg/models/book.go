package models

//Book struct
type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	ISBN   string `json:"isbn"`
	Type   string `json:"type"`
	Author Author `json:"author"`
}
