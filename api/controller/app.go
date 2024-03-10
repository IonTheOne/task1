package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Mlstermass/task1/pkg/env"
	"github.com/Mlstermass/task1/storage"
)

const (
	NewsItemIDStr = "newsItemId"
)

type App struct {
	config  env.Config
	storage storage.DocumentActions
}

func NewApp(
	config env.Config,
	storage storage.DocumentActions,
) App {
	return App{
		config:  config,
		storage: storage,
	}
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// @Summary Get all news
// @Description Fetch all news items from the database
// @Tags news
// @Produce json
// @Success 200 {array} httpentity.NewsItem
// @Router /news [get]
func (a *App) GetNews(w http.ResponseWriter, r *http.Request) {
	newsItems, err := a.storage.GetNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newsItems)
}

// @Summary Get news by ID
// @Description Fetch a single news item from the database by ID
// @Tags news
// @Produce json
// @Param newsItemId path string true "News Item ID"
// @Success 200 {object} httpentity.NewsItem
// @Router /news/{newsItemId} [get]
func (a *App) GetNewsByID(w http.ResponseWriter, r *http.Request) {
	newsItemId := chi.URLParam(r, NewsItemIDStr)
	newsItem, err := a.storage.GetNewsByID(newsItemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newsItem)
}
