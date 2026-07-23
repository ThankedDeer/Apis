package handlers

import (
	httpclient "apis/http_client"
	"apis/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	result := models.Result[models.HealthStatus]{
		Data: models.HealthStatus{
			Status:  "ok",
			Message: "Hello, server healthy",
		},
	}
	c.JSON(http.StatusOK, result)
}

func Dashboard(c *gin.Context) {
	ctx := c.Request.Context()
	var wg sync.WaitGroup
	client := httpclient.Newclient()
	respCha := make(chan models.TaskResult, 5)

	response := &models.Result[models.DashboardData]{
		Data: models.DashboardData{
			NewsList:   []models.News{},
			MusicList:  models.Music{},
			Canonical:  models.Canonical{},
			Breweries:  []models.Breweries{},
			WhaterList: []models.Whater{},
		},
		Error:   nil,
		Message: "",
		Code:    0,
	}

	wg.Add(5)
	go getNews(ctx, client, respCha, &wg)
	go getMusic(ctx, client, respCha, &wg)
	go getCanonical(ctx, client, respCha, &wg)
	go getBreweries(ctx, client, respCha, &wg)
	go getWhater(ctx, client, respCha, &wg)

	go func() {
		wg.Wait()
		close(respCha)
	}()

	errors := make(map[string]error)

	for task := range respCha {
		if task.Error != nil {
			errors[task.Source] = task.Error
			continue
		}

		switch task.Source {
		case "news":
			if news, ok := task.Data.([]models.News); ok {
				response.Data.NewsList = news
			}
		case "music":
			if music, ok := task.Data.(models.Music); ok {
				response.Data.MusicList = music
			}
		case "canonical":
			if canonical, ok := task.Data.(models.Canonical); ok {
				response.Data.Canonical = canonical
			}
		case "breweries":
			if breweries, ok := task.Data.([]models.Breweries); ok {
				response.Data.Breweries = breweries
			}
		case "whater":
			if whaterList, ok := task.Data.([]models.Whater); ok {
				response.Data.WhaterList = whaterList
			}
		}
	}

	if len(errors) > 0 {
		response.Message = "partial data available"
		response.Code = http.StatusOK
	} else {
		response.Message = "All data retrieved successfully"
		response.Code = http.StatusOK
	}

	if len(errors) > 0 {
		for src, err := range errors {
			log.Printf("Error in %s: %v", src, err)
		}
	}
	c.JSON(http.StatusOK, response)
}

func getNews(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
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

func getMusic(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
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

func getCanonical(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
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

func getBreweries(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
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

func getWhater(ctx context.Context, client *httpclient.Client, ch chan<- models.TaskResult, wg *sync.WaitGroup) {
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
