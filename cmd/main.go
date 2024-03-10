package main

import (
	"context"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mlstermass/task1/api/controller"
	"github.com/Mlstermass/task1/api/controller/httpentity"
	"github.com/Mlstermass/task1/api/router"
	"github.com/Mlstermass/task1/pkg/env"
	"github.com/Mlstermass/task1/storage"
	"github.com/Mlstermass/task1/storage/mongodb"
)

const (
	newsServiceURL = "https://www.htafc.com/api/incrowd/getnewlistinformation?count=50"
)

// @title			Task1 API
// @version		1.0
// @description	API for task1 service
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host		localhost:8100
// @BasePath	/
func main() {
	// load env variables to the Config struct
	var conf env.Config
	env.LoadConfig(&conf)

	listeningAddress := flag.String("listening-address",
		conf.AppHost, "Address which server handle")
	flag.Parse()

	var mongoConn mongodb.ConnMongo
	mongoDBConn, err := mongoConn.NewConnMongo(conf)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("Successfully connected to MongoDB")

	defer func() {
		err = mongoDBConn.Disconnect(context.Background())
		if err != nil {
			log.Println("Failed to stop mongodb connector")
		}
		log.Println("Teared down mongodb")
	}()

	mongoDriver := mongodb.NewMongo(mongoDBConn, conf)

	ticker := time.NewTicker(conf.RefreshInterval)

	// Add news articles only on the first run
	_, err = os.Stat("news_seeded")
	if os.IsNotExist(err) {
		// Add news articles to the database on startup
		allNews, err := getNews(newsServiceURL)
		if err != nil {
			log.Fatalf("Failed to get news: %v", err)
		}
		for _, newsItem := range allNews.NewsItems {
			err = mongoDriver.AddNews(newsItem)
			if err != nil {
				log.Printf("Failed to add news: %v", err)
			}
		}
		log.Println("Successfully added news")

		// Create the file to indicate that the news has been added
		_, err = os.Create("news_seeded")
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
	}

	// Update news articles by interval
	go func() {
		for range ticker.C {
			news, err := getNews(newsServiceURL)
			if err != nil {
				log.Printf("Failed to get news: %v", err)
				continue
			}

			for _, newsItem := range news.NewsItems {
				exists, err := mongoDriver.NewsExists(newsItem.NewsArticleID)
				if err != nil {
					log.Printf("Failed to check if news exists: %v", err)
					continue
				}
				if exists {
					log.Printf("News with ID %s already exists, skipping", newsItem.NewsArticleID)
					continue
				}

				err = mongoDriver.AddNews(newsItem)
				if err != nil {
					log.Printf("Failed to add news: %v", err)
					continue
				}
				log.Printf("Successfully added news: %v", newsItem)
			}
		}
	}()

	ctl := newControllers(conf, mongoDriver)

	r := router.New(ctl, conf)

	srv := &http.Server{
		Handler: r,
		Addr:    *listeningAddress,
		// Enforce timeouts for servers you create!
		WriteTimeout: conf.WriteTimeout,
		ReadTimeout:  conf.ReadTimeout,
	}
	log.Println("Listening to", *listeningAddress)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func newControllers(
	config env.Config,
	mongoDriver storage.DocumentActions,
) controller.App {
	return controller.NewApp(
		config, mongoDriver)
}

func getNews(newsServiceURL string) (httpentity.News, error) {
	// Make the GET request
	resp, err := http.Get(newsServiceURL)
	if err != nil {
		return httpentity.News{}, fmt.Errorf("Failed to make GET request to news service: %v", err)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return httpentity.News{}, fmt.Errorf("Failed to read response body: %v", err)
	}

	var news httpentity.News
	err = xml.Unmarshal(body, &news)
	if err != nil {
		return httpentity.News{}, fmt.Errorf("Failed to unmarshal XML: %v", err)
	}

	return news, nil
}
