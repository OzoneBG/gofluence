package models

// Article represent a single article in the application
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Articles represent a collection of articles in the application
type Articles []Article
