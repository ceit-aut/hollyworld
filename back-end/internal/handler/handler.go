package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// Handler is our endpoint handling struct
type Handler struct {
	Db *sql.DB
}

// GetTopMovies returns top 6 movies
func (h *Handler) GetTopMovies(c *fiber.Ctx) error {
	// creating a movie type
	type Movie struct {
		Uid    int    `json:"uid"`
		Title  string `json:"title"`
		Poster string `json:"poster"`
	}

	// creating our variables
	var (
		movie  Movie
		movies []Movie

		query = "SELECT uid, title, poster FROM movies ORDER BY score LIMIT 6"
	)

	// executing database query
	rows, err := h.Db.Query(query)
	if err != nil {
		return err
	}

	// scan output
	for rows.Next() {
		er := rows.Scan(&movie)
		if er != nil {
			return er
		}

		movies = append(movies, movie)
	}

	// close rows
	_ = rows.Close()

	return c.JSON(movies)
}

// GetSingleMovie returns a single movie information by id
func (h *Handler) GetSingleMovie(c *fiber.Ctx) error {
	return nil
}

// GetMovieFile returns a movie thriller file
func (h *Handler) GetMovieFile(c *fiber.Ctx) error {
	return nil
}
