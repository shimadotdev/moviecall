package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKeyword(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"./moviecall", "-m", "atlas"}

	movie, tv, trendingMovies, trendingTvs := GetKeyword()

	assert.Equal(t, "atlas", movie, "Expected movie name to be 'atlas'")
	assert.Equal(t, "", tv, "Expected TV show name to be empty")
	assert.False(t, trendingMovies, "Expected trendingMovies to be false")
	assert.False(t, trendingTvs, "Expected trendingTvs to be false")

}
