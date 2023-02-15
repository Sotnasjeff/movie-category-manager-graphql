package graph

import "github.com/Sotnasjeff/movie-category-manager-graphql/internal/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *db.Category
	MovieDB    *db.Movie
}
