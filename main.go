// Recipes API
//
// This is a recipes API.
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//    Schemes: http
//    Host: localhost:8080
//    BasePath: /
//    Version: 0.0.1
//    License: MIT http://opensource.org/licenses/MIT
//    Contact: Francisco Oliveira<delley.fx@gmail.com> https://delley.github.io/
//
//    Consumes:
//    - application/json
//
//    Produces:
//    - application/json
// swagger:meta
package main

import (
	"context"
	"log"
	"os"

	handlers "github.com/delley/recipes-api/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var collection *mongo.Collection
var recipesHandler *handlers.RecipesHandler

func init() {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	recipesHandler = handlers.NewRecipesHandler(ctx, collection)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	router.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	router.Run()
}
