package tmdb

import (
	"fmt"

	app "github.com/shimadotdev/moviecall/internal"
)

type MovieDataCollection struct {
	Page         int         `json:"page"`
	Results      []MovieData `json:"results"`
	TotalPages   int         `json:"total_pages"`
	TotalResults int         `json:"total_results"`
}

type MovieData struct {
	Id          int32   `json:"id"`
	Language    string  `json:"original_language"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	Rating      float32 `json:"vote_average"`
	ReleaseDate string  `json:"release_date"`
	Genres      []int   `json:"genre_ids"`
}

type MovieDataDetail struct {
	Adult       bool    `json:"adult"`
	Id          int64   `json:"id"`
	Language    string  `json:"original_language"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	Rating      float32 `json:"vote_average"`
	ReleaseDate string  `json:"release_date"`
	Genres      []struct {
		Name string
	}
	PosterPath string `json:"poster_path"`
	Homepage   string `json:"homepage"`
}

func formatMovieDataDetail(detail MovieDataDetail) []string {
	year := ""
	if len(detail.ReleaseDate) >= 4 {
		year = detail.ReleaseDate[:4]
	}
	url := fmt.Sprintf("https://www.themoviedb.org/movie/%d-%s", detail.Id, app.ConvertString(detail.Title))

	return []string{
		app.EllipsizeString(detail.Title, 30),
		year,
		detail.Language,
		fmt.Sprintf("%.1f", detail.Rating),
		app.FormatGenres(detail.Genres),
		url,
	}
}
