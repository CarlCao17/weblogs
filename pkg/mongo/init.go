package mongo

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/weblogs/pkg/log"
)

func InitMongo() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(GetBaseURI()).SetServerAPIOptions(serverAPI)

	var err error
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err = client.Database("test").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	log.Infof("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}

func GetBaseURI() string {
	baseURI := os.Getenv("MONGODB_URL")
	appName := os.Getenv("APP_NAME")
	if appName != "" {
		baseURI = baseURI + "&appName=" + appName
	}
	return baseURI
}
