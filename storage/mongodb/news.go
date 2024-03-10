package mongodb

import (
	"context"

	"github.com/Mlstermass/task1/api/controller/httpentity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Mongo) UseNews() *mongo.Collection {
	return m.client.Database(m.conf.AppMongoDBName).Collection(m.conf.AppMongoCollectionName)
}

func (m *Mongo) NewsExists(newsArticleID string) (bool, error) {
	filter := bson.M{"newsArticleID": newsArticleID}
	err := m.UseNews().FindOne(context.Background(), filter).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *Mongo) AddNews(newsItem httpentity.NewsItem) error {
	_, err := m.UseNews().InsertOne(context.Background(), newsItem)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) GetNews() ([]httpentity.NewsItem, error) {
	var newsItems []httpentity.NewsItem
	cursor, err := m.UseNews().Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var newsItem httpentity.NewsItem
		err := cursor.Decode(&newsItem)
		if err != nil {
			return nil, err
		}
		newsItems = append(newsItems, newsItem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return newsItems, nil
}

func (m *Mongo) GetNewsByID(newsItemID string) (httpentity.NewsItem, error) {
	filter := bson.M{"newsArticleID": newsItemID}
	var newsItem httpentity.NewsItem
	err := m.UseNews().FindOne(context.Background(), filter).Decode(&newsItem)
	if err != nil {
		return httpentity.NewsItem{}, err
	}
	return newsItem, nil
}
