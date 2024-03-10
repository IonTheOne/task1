package storage

import (
	"fmt"

	"github.com/Mlstermass/task1/api/controller/httpentity"
	"github.com/stretchr/testify/mock"
)

type MockDocumentActions struct {
	mock.Mock
}

func (m *MockDocumentActions) NewsExists(newsArticleID string) (bool, error) {
	args := m.Called(newsArticleID)
	return args.Bool(0), args.Error(1)
}

func (m *MockDocumentActions) AddNews(newsItem httpentity.NewsItem) error {
	args := m.Called(newsItem)
	return args.Error(0)
}

func (m *MockDocumentActions) GetNews() ([]httpentity.NewsItem, error) {
	args := m.Called()
	return args.Get(0).([]httpentity.NewsItem), args.Error(1)
}

func (m *MockDocumentActions) GetNewsByID(newsItemID string) (httpentity.NewsItem, error) {
	if newsItemID == "" {
		return httpentity.NewsItem{}, fmt.Errorf("newsItemID cannot be empty")
	}
	args := m.Called(newsItemID)
	return args.Get(0).(httpentity.NewsItem), args.Error(1)
}
