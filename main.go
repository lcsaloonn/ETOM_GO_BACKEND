package main

import (
	"api/ETOM/albums/controllers"
	"api/ETOM/albums/services"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	albumService services.AlbumService
	AlbumController controllers.AlbumController
	ctx context.Context
	albumsCollection *mongo.Collection
	mongoClient *mongo.Client
	err error
)

func init(){
	ctx = context.TODO()
	mongoConnection := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongDB")
	albumsCollection = mongoClient.Database("albumDb").Collection("Albums")
	albumService = services.NewAlbumService(albumsCollection, ctx) 
	AlbumController = controllers.New(albumService)
	server = gin.Default()

}

func main() {
	defer mongoClient.Disconnect(ctx)
	basePath := server.Group("v1/")
	AlbumController.RegisterAlbumRoutes(basePath)
	log.Fatal(server.Run(":9090"))

}