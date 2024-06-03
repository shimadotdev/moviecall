package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	app "github.com/shimadotdev/moviecall/internal"
)

type ConfigFile struct {
	ApiKey     string `json:"api_key"`
	ApiBaseUrl string `json:"api_base_url"`
}

type TMDB struct {
	ApiKey     string
	ApiBaseUrl string
}

func Initiate() (*TMDB, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting working directory: %v", err)
	}

	relativePath := filepath.Join(wd, "../config.json")
	configData, err := os.ReadFile(relativePath)
	if err != nil {
		return nil, fmt.Errorf("error reading configuration file: %v", err)
	}

	var configFile ConfigFile
	if err := json.Unmarshal(configData, &configFile); err != nil {
		return nil, fmt.Errorf("error unmarshaling configuration file: %v", err)
	}

	return &TMDB{
		ApiKey:     configFile.ApiKey,
		ApiBaseUrl: configFile.ApiBaseUrl,
	}, nil
}

func getResult[T any](endpoint string, payload T) (T, error) {
	response, err := http.Get(endpoint)
	if err != nil {
		return payload, fmt.Errorf("error requesting API endpoint: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return payload, fmt.Errorf("error reading API response: %v", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return payload, fmt.Errorf("HTTP error (code %d): %s", response.StatusCode, response.Status)
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return payload, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return payload, nil
}


func SearchByKeyword(searchType, searchTerm string) error {

	tmdb, err := Initiate()
	if err != nil {
		return err
	}

	var (
		list       [][]string
		header     []string
		tableTitle string
	)

	switch searchType {
	case "tv":
		payload := tmdb.searchTv(searchTerm)
		idList := getIdListFromPayload(payload.Results)
		list, err = getDetailsByIdList("tv", idList)
		if err != nil {
			return err
		}
		tableTitle = "Results for TV show: " + searchTerm
		header = []string{"Title", "First Air Date", "Language", "Rating", "Genres", "Link"}
	case "movie":
		payload := tmdb.searchMovie(searchTerm)
		idList := getIdListFromPayload(payload.Results)
		list, err = getDetailsByIdList("movie", idList)
		if err != nil {
			return err
		}
		tableTitle = "Results for movie: " + searchTerm
		header = []string{"Title", "Release Date", "Language", "Rating", "Genres",  "Link"}
	default:
		return fmt.Errorf("invalid search type: %s", searchType)
	}

	app.PrintTable(header, list, tableTitle)
	return nil
}

func getIdListFromPayload[T any](results []T) []int32 {
	var idList []int32
	for _, item := range results {
		switch v := any(item).(type) {
		case TvData:
			idList = append(idList, v.Id)
		case MovieData:
			idList = append(idList, v.Id)
		}
	}
	return idList
}

func getDetailsByIdList(searchType string, idList []int32) ([][]string, error) {
	var wg sync.WaitGroup
	resultCh := make(chan any, len(idList))

	tmdb, err := Initiate()
	if err != nil {
		return nil, err
	}

	for _, id := range idList {
		wg.Add(1)
		go func(id int32) {
			defer wg.Done()
			switch searchType {
			case "tv":
				res := tmdb.getTvById(id)
				resultCh <- res
			case "movie":
				res := tmdb.getMovieById(id)
				resultCh <- res
			}
		}(id)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var resultList [][]string
	for result := range resultCh {
		switch v := result.(type) {
		case TvDataDetail:
			resultList = append(resultList, formatTvDataDetail(v))
		case MovieDataDetail:
			resultList = append(resultList, formatMovieDataDetail(v))
		}
	}
	return resultList, nil
}

func (t *TMDB) getMovieById(id int32) MovieDataDetail {
	endpoint := fmt.Sprintf("%s/movie/%d?api_key=%s", t.ApiBaseUrl, id, t.ApiKey)
	var movieDetail MovieDataDetail
	res, err := getResult(endpoint, movieDetail)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (t *TMDB) getTvById(id int32) TvDataDetail {
	endpoint := fmt.Sprintf("%s/tv/%d?api_key=%s", t.ApiBaseUrl, id, t.ApiKey)
	var tvDetailData TvDataDetail
	res, err := getResult(endpoint, tvDetailData)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (t *TMDB) searchTv(keyword string) TvDataCollection {

	endpoint := fmt.Sprintf("%s/search/tv?query=%s&api_key=%s", t.ApiBaseUrl, app.ConvertString(keyword), t.ApiKey)

	var tvCollection TvDataCollection

	res, err := getResult(endpoint, tvCollection)

	if err != nil {
		log.Fatal(err)
	}

	return res

}

func (t *TMDB) searchMovie(keyword string) MovieDataCollection {

	endpoint := fmt.Sprintf("%s/search/movie?query=%s&api_key=%s", t.ApiBaseUrl, app.ConvertString(keyword), t.ApiKey)
	var movieCollection MovieDataCollection

	res, err := getResult(endpoint, movieCollection)

	if err != nil {
		log.Fatal(err)
	}

	return res

}

