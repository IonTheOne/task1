package mongodb

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Mlstermass/task1/pkg/env"
	"github.com/Mlstermass/task1/storage"
)

const (
	mongoDBInitErrStr = "unable to initialize connection"
	mongoDBConnErrStr = "unable to connect to the mongodb db"
)

type ConnMongo struct{}

func (m *ConnMongo) NewConnMongo(
	conf env.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	clientOptions := options.Client().ApplyURI(conf.AppMongoDBConnectionString).SetDirect(true)
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(mongoDBInitErrStr, err)
	}
	err = c.Ping(ctx, nil)
	if err != nil {
		log.Fatal(mongoDBConnErrStr, err)
	}
	return c, err
}

type Mongo struct {
	client *mongo.Client
	conf   env.Config
}

var once sync.Once

func NewMongo(client *mongo.Client, conf env.Config) storage.DocumentActions {
	dbName := conf.AppMongoDBName
	collectionName := conf.AppMongoCollectionName
	once.Do(func() {
		_ = client.Database(dbName).Collection(collectionName)
		_ = client.Database(dbName).CreateCollection(context.Background(), collectionName)
	})
	return &Mongo{client: client, conf: conf}
}
