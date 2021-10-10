package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jiel/golden/controllers"
	daoimpl "github.com/jiel/golden/dao/impl"
	serviceimpl "github.com/jiel/golden/services/impl"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUri = flag.String("mongo-uri", "", "")
var databaseName = flag.String("database-name", "", "")
var albumCollectionName = flag.String("album-collection-name", "", "")

func main() {
	flag.Parse()
	
	opts := options.Client().ApplyURI(*mongoUri)
	client, err := mongo.NewClient(opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := client.Connect(context.Background()); err != nil {
		fmt.Println(err)
		return
	}

	albumRepo := daoimpl.NewMongoAlbumRepo(client, *databaseName, *albumCollectionName)
	albumService := serviceimpl.NewAlbumService(albumRepo)
	albumController := controllers.NewAlbumController(albumService)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	v1 := r.Group("/api/v1")
	v1.GET("/albums", albumController.GetAlbums)
	v1.POST("/albums/create", albumController.CreateAlbum)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
