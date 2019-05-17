package models

import jwt "github.com/dgrijalva/jwt-go"

// Article represent a single article in the application.
type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"`
}

// Articles represent a collection of articles in the application.
type Articles []Article

// Token represents the user token
type Token struct {
	UserID int
	jwt.StandardClaims
}
