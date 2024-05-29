package app

import (
	"flag"
)

func GetKeyword() (string, string) {

	var movie string
	var tv string

	flag.StringVar(&movie, "m", "", "name of the movie you are looking for!")
	flag.StringVar(&tv, "t", "", "name of the TV show you are looking for!")
	flag.Parse()
	return movie, tv
}
