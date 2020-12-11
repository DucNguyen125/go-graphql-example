package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example/routers"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Register Router
	routersInit := routers.InitRouter()
	if err != nil {
		log.Fatal("Error get PORT")
	}
	endPoint := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	if err = server.ListenAndServe(); err != nil {
		log.Fatal("Fail to start error server")
	}
}
