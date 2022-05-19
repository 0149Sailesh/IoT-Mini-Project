package services

import(
	"context"
	"time"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/0149Sailesh/iot-server/config"
)
var DB *mongo.Client 

func ConnectInit(){
    DB = ConnectDB()
}

func ConnectDB() *mongo.Client  {
	var l = log.WithFields(log.Fields{
		"method":"ConnectDB",
	})
    client, err := mongo.NewClient(options.Client().ApplyURI(config.EnvMongoURI()))
    if err != nil {
        l.Fatal(err)
    }
  
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
		panic("Terminating program")
    }
    l.Info("Connected to MongoDB")
    return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("IoT").Collection(collectionName)
    return collection
}