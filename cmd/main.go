package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Mlstermass/task1/api/controller"
	"github.com/Mlstermass/task1/api/router"
	"github.com/Mlstermass/task1/pkg/env"
	"github.com/Mlstermass/task1/storage"
	"github.com/Mlstermass/task1/storage/mongodb"
)

// @title			Documents
// @version		1.0
// @description	API for documents service
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host		localhost:8000
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
		os.Exit(1)
	}
	log.Println("Successfully connected to MongoDB")

	defer func(mongoDBConn *mongo.Client) {
		err = mongoDBConn.Disconnect(context.Background())
		if err != nil {
			log.Println("Failed to stop mongodb connector")
		}
		log.Println("Teared down mongodb")
	}(mongoDBConn)

	mongoDriver := mongodb.NewMongo(mongoDBConn, conf)

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
		return
	}
}

func newControllers(
	config env.Config,
	mongoDriver storage.DocumentActions,
) controller.App {
	return controller.NewApp(
		config, mongoDriver)
}
