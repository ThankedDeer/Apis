package services

import (
	httpclient "apis/http_client"
	"apis/models"
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

func GetNews(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()

	const source = "news"

	resp, err := client.CreateRequest(
		ctx,
		"https://juliayxhuang.github.io/florida-man-api/api/headlines.json",
		http.MethodGet,
		nil,
		map[string]string{},
	)
	if err != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   []models.News{},
			Error:  err,
		}
		return
	}
	defer resp.Body.Close()
	var news []models.News
	if decErr := json.NewDecoder(resp.Body).Decode(&news); decErr != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   []models.News{},
			Error:  decErr,
		}
		return
	}
	ch <- models.TaskResult{
		Source: source,
		Data:   news,
		Error:  nil,
	}
}

func GetMusic(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	const source = "music"
	resp, err := client.CreateRequest(
		ctx,
		"https://rest.bandsintown.com/artists/Skrillex?app_id=67",
		http.MethodGet,
		nil,
		map[string]string{},
	)
	if err != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   models.Music{},
			Error:  err,
		}
		return
	}
	defer resp.Body.Close()

	var music models.Music
	if decErr := json.NewDecoder(resp.Body).Decode(&music); decErr != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   models.Music{},
			Error:  decErr,
		}
		return
	}

	ch <- models.TaskResult{
		Source: source,
		Data:   music,
		Error:  nil,
	}
}

func GetCanonical(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	const source = "canonical"

	resp, err := client.CreateRequest(
		ctx,
		"https://alpha.moss.land/api/canonical/entities.json",
		http.MethodGet,
		nil,
		map[string]string{},
	)
	if err != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   models.Canonical{},
			Error:  err,
		}
		return
	}
	defer resp.Body.Close()

	var canonical models.Canonical
	if decErr := json.NewDecoder(resp.Body).Decode(&canonical); decErr != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   models.Canonical{},
			Error:  decErr,
		}
		return
	}
	ch <- models.TaskResult{
		Source: source,
		Data:   canonical,
		Error:  nil,
	}
}

func GetBreweries(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	const source = "breweries"

	resp, err := client.CreateRequest(
		ctx,
		"https://api.openbrewerydb.org/v1/breweries/",
		http.MethodGet,
		nil,
		map[string]string{},
	)
	if err != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   []models.Breweries{},
			Error:  err,
		}
		return
	}
	defer resp.Body.Close()

	var breweries []models.Breweries
	if decErr := json.NewDecoder(resp.Body).Decode(&breweries); decErr != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   []models.Breweries{},
			Error:  decErr,
		}
		return
	}
	ch <- models.TaskResult{
		Source: source,
		Data:   breweries,
		Error:  nil,
	}
}

func GetWhater(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	const source = "whater"

	resp, err := client.CreateRequest(
		ctx,
		"https://aviationweather.gov/api/data/metar?ids=KJFK,KLAX,KMIA,KORD,KDFW,KDEN,KSFO,KSEA&format=json",
		http.MethodGet,
		nil,
		map[string]string{},
	)
	if err != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   []models.Whater{},
			Error:  err,
		}
		return
	}
	defer resp.Body.Close()

	var whater []models.Whater
	if decErr := json.NewDecoder(resp.Body).Decode(&whater); decErr != nil {
		ch <- models.TaskResult{
			Source: source,
			Data:   []models.Whater{},
			Error:  decErr,
		}
		return
	}
	ch <- models.TaskResult{
		Source: source,
		Data:   whater,
		Error:  nil,
	}
}
