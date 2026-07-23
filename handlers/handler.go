package handlers

import (
	httpclient "apis/http_client"
	"apis/models"
	"apis/services"
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
	go services.GetNews(ctx, client, respCha, &wg)
	go services.GetMusic(ctx, client, respCha, &wg)
	go services.GetCanonical(ctx, client, respCha, &wg)
	go services.GetBreweries(ctx, client, respCha, &wg)
	go services.GetWhater(ctx, client, respCha, &wg)

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
