package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mlstermass/task1/api/controller/httpentity"
	"github.com/Mlstermass/task1/pkg/env"
	"github.com/Mlstermass/task1/storage"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestApp_GetNews(t *testing.T) {
	mockStorage := &storage.MockDocumentActions{}
	mockStorage.On("GetNews").Return([]httpentity.NewsItem{}, nil)

	app := NewApp(env.Config{}, mockStorage)

	req, err := http.NewRequest("GET", "/news", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	app.GetNews(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	mockStorage.AssertCalled(t, "GetNews")
}

func TestApp_GetNewsByID(t *testing.T) {
	newsItemID := "65401"
	newsItem := httpentity.NewsItem{}

	mockStorage := &storage.MockDocumentActions{}
	mockStorage.On("GetNewsByID", newsItemID).Return(newsItem, nil)

	app := NewApp(env.Config{}, mockStorage)

	req, err := http.NewRequest("GET", "/news/"+newsItemID, nil)
	assert.NoError(t, err)

	// Set the URL parameter in the request context
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(NewsItemIDStr, newsItemID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	recorder := httptest.NewRecorder()

	app.GetNewsByID(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	// Check the response body
	var responseNewsItem httpentity.NewsItem
	err = json.Unmarshal(recorder.Body.Bytes(), &responseNewsItem)
	assert.Equal(t, newsItem, responseNewsItem)

	mockStorage.AssertCalled(t, "GetNewsByID", newsItemID)
}
