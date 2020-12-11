package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example/routers"
	"example/utils/mysql_util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(os.Getenv("RUN_MODE"))

	// Connect to database
	if err = mysql_util.Connect(); err != nil {
		log.Fatal("Error connecting to database")
	}

	// Migrate database
	if err = mysql_util.AutoMigrate(); err != nil {
		log.Fatal("Error migrating to database")
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
