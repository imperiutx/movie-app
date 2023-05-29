package repository

import (
	"database/sql"
	"movie-app-be/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
