package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Movie struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	Category    string
}

func NewMovie(db *sql.DB) *Movie {
	return &Movie{db: db}
}

func (m *Movie) CreateMovie(name string, description string, categoryId string) (Movie, error) {
	id := uuid.New().String()

	_, err := m.db.Exec("INSERT INTO movies (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryId)
	if err != nil {
		return Movie{}, err
	}
	return Movie{ID: id, Name: name, Description: description}, nil
}

func (m *Movie) FindAllMovies() ([]Movie, error) {
	rows, err := m.db.Query("SELECT id, name, description, category_id FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	movies := []Movie{}
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}
		movies = append(movies, Movie{ID: id, Name: name, Description: description, Category: categoryId})
	}
	return movies, nil
}

func (m *Movie) FindByCategoryID(categoryID string) ([]Movie, error) {
	rows, err := m.db.Query("SELECT id, name, description, category_id FROM movies WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	movies := []Movie{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		movies = append(movies, Movie{
			ID:          id,
			Name:        name,
			Description: description,
			Category:    categoryID,
		})
	}

	return movies, nil
}
