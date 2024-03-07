package controller

import (
	"net/http"

	"github.com/Mlstermass/task1/pkg/env"
	"github.com/Mlstermass/task1/storage"
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

func (a *App) GetNews(w http.ResponseWriter, r *http.Request) {
	// some code
	w.Write([]byte("news"))
}
