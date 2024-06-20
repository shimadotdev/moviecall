package main

import (
	"fmt"
	"os"

	app "github.com/shimadotdev/moviecall/internal"
	"github.com/shimadotdev/moviecall/internal/tmdb"
)

func main() {
	movie, tv, trendingMovies, trendingTvs := app.GetKeyword()

	var category, subject string
	switch {
		case movie != "":
			category, subject = "movie", movie
		case tv != "":
			category, subject = "tv", tv
		case trendingMovies:
			category, subject = "trendingMovies", ""
		case trendingTvs:
			category, subject = "trendingTvs", ""
		default:
			fmt.Fprintln(os.Stderr, "Error: Please specify either a movie or TV show.")
			os.Exit(1)
	}

	if err := tmdb.SearchByKeyword(category, subject); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
