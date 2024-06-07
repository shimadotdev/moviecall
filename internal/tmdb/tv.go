package tmdb

import (
	"fmt"

	app "github.com/shimadotdev/moviecall/internal"
)

type TvDataCollection struct {
	Page         int      `json:"page"`
	Results      []TvData `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

type TvData struct {
	Id           int32   `json:"id"`
	Language     string  `json:"original_language"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	Rating       float32 `json:"vote_average"`
	FirstAirDate string  `json:"first_air_date"`
	Genres       []int   `json:"genre_ids"`
}

type TvDataDetail struct {
	Adult        bool    `json:"adult"`
	Id           int64   `json:"id"`
	Language     string  `json:"original_language"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	Rating       float32 `json:"vote_average"`
	FirstAirDate string  `json:"first_air_date"`
	Genres       []struct {
		Name string
	}
	PosterPath string `json:"poster_path"`
	Homepage   string `json:"homepage"`
}

func formatTvDataDetail(detail TvDataDetail) []string {
	year := ""
	if len(detail.FirstAirDate) >= 4 {
		year = detail.FirstAirDate[:4]
	}

	url := fmt.Sprintf("https://www.themoviedb.org/tv/%d", detail.Id)

	return []string{
		app.EllipsizeString(detail.Name, 30),
		year,
		detail.Language,
		fmt.Sprintf("%.1f", detail.Rating),
		app.FormatGenres(detail.Genres),
		url,
	}
}
