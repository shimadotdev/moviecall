package app

import (
	"flag"
)

func GetKeyword() (string, string, bool, bool) {

	var movie string
	var tv string
	var trendingMovies bool
	var trendingTvs bool

	flag.StringVar(&movie, "m", "", "name of the movie you are looking for!")
	flag.StringVar(&tv, "t", "", "name of the TV show you are looking for!")
	flag.BoolVar(&trendingMovies, "tm", false, "Trendings movies")
	flag.BoolVar(&trendingTvs, "tt", false, "Trendings Tv shows")
	flag.Parse()
	return movie, tv, trendingMovies, trendingTvs
}
