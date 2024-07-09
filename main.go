package main

import (
	_ "example/blog/docs"
	"example/blog/initializers"
	"example/blog/router"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func init() {
	if os.Getenv("DEPLOYMENT_TYPE") != "k8s" {
		initializers.LoadEnvVariables()
	}
	initializers.ConnectToDB()
}

// @title 	Blog Service API
// @version	1.0
// @description A Blog Service API in Go using Gin framework

// @host 	localhost:3000
// @BasePath /
// @schemes http

// @securityDefinitions.apiKey JWT
// @in header
// @name token
func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.Info().Msg("Started Server!")

	router := router.NewRouter()
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal().Msg("Server cannot start serve the Blog API")
	}
}
