package routes

import (
	"fmt"
	"net/http"

	"github.com/ozonebg/gofluence/context"
)

// Route is a representation of a api defined route.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of api defined routes.
type Routes []Route

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "REST API: v0.1.0")
}

// GetRoutes returns all api-defined routes.
func GetRoutes(context *context.Context) Routes {
	routes := Routes{
		// Index
		Route{"Index", "GET", "/", index},

		// Articles
		Route{"CreateArticle", "POST", "/api/article", context.ArticlesController.CreateArticle},
		Route{"AllArticles", "GET", "/api/articles", context.ArticlesController.AllArticles},
		Route{"GetArticle", "GET", "/api/article/{id}", context.ArticlesController.GetArticle},
		Route{"UpdateArticle", "PUT", "/api/article/{id}", context.ArticlesController.UpdateArticle},
		Route{"DeleteArticle", "DELETE", "/api/article/{id}", context.ArticlesController.DeleteArticle},

		// Users
		Route{"CreateUser", "POST", "/api/user/new", context.UsersController.CreateUser},
		Route{"AllUsers", "GET", "/api/users", context.UsersController.AllUsers},
		Route{"GetUser", "GET", "/api/user/{id}", context.UsersController.GetUser},
		Route{"UpdateUser", "PUT", "/api/user/{id}", context.UsersController.UpdateUser},
		Route{"DeleteUser", "DELETE", "/api/user/{id}", context.UsersController.DeleteUser},
		Route{"Authenticate", "POST", "/api/login", context.UsersController.Authenticate},
	}

	return routes
}
