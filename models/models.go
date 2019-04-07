package models

// Article represent a single article in the application.
type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"`
}

// Articles represent a collection of articles in the application.
type Articles []Article

// User represents a single user in the application.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Users represent a collection of users in the application.
type Users []User
