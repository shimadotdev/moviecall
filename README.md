[![Go Reference](https://pkg.go.dev/badge/github.com/shimadotdev/moviecall.svg)](https://pkg.go.dev/github.com/shimadotdev/moviecall)
[![Tests](https://github.com/shimadotdev/moviecall/actions/workflows/test.yml/badge.svg)](https://github.com/shimadotdev/moviecall/actions/workflows/test.yml)

# Moviecall

## Overview

**Moviecall** is a CLI application designed to search for movies and TV shows. It is a learning project to practice and improve skills in Go, specifically focusing on concurrency and goroutines. The project will be updated over time to include more features and improvements.

## Features
- List of Trending movies and TV show
- Search for movies by title
- Search for TV shows by title

## Installation

To run the application, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the [official Go website](https://golang.org/dl/).

## TMDB API

The application uses the TMDB (The Movie Database) API for retrieving movie and TV show data. You can sign up for a TMDB API key by visiting [TMDB website](https://developer.themoviedb.org/) and following the instructions.

## Build

To build the application, ensure you have created a .env file based on .env.example and added your TMDB API key. Then, run:

```bash
make
```
By running this command, the binary executable file will be placed in the bin directory. From there, you can execute the commands provided by the application.

## Clean

To remove the binary execution file, run:

```bash
make clean
```

## Usage

To use the application, you can utilize the commands below:

### Trending Movies
To fetch a list of trending movies right now use the `-tm` flag

```bash
./moviecall -tm
```

### Trending TV shows
To fetch a list of trending movies right now use the `-tt` flag

```bash
./moviecall -tt
```

### Search for a Movie

To search for a movie by its title, use the `-m` flag followed by the movie name:

```bash
./moviecall -m "movie name"
```

### Search for a TV Show

To search for a TV show by its title, use the -t flag followed by the TV show name:

```bash
./moviecall -t "tvshow name"
```

### Contributing
Contributions are welcome! If you have suggestions for improvements or find bugs, please open an issue or submit a pull request.

### License
This package is open-source software licensed under the [MIT licensev](https://opensource.org/licenses/MIT).

### Contact
For any questions or feedback, please contact [hi@shima.dev](mailto:hi@shima.dev)

