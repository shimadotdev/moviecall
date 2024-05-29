package main

import (
	"fmt"
	"os"

	app "github.com/shimadotdev/moviecall/internal"
	"github.com/shimadotdev/moviecall/internal/tmdb"
)

func main() {
	movie, tv := app.GetKeyword()

	if movie != "" {
		if err := tmdb.SearchByKeyword("movie", movie); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else if tv != "" {
		if err := tmdb.SearchByKeyword("tv", tv); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Error: Please specify either a movie or TV show.")
		os.Exit(1)
	}
}
