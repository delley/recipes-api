package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name        string     `json:"name"`
	Tags        []string   `json:"tags"`
	Ingredients []string   `json:"ingredients"`
	Intructions []string   `json:"instructions"`
	PublishedAt time.Timer `json:"published_at"`
}

func main() {
	router := gin.Default()
	router.Run()
}
