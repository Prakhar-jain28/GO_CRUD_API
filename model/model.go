package model

type Blog struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Content string `json:"content"`
}