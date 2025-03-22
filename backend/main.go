package main

import (
	"leal/cmd/api"
	_ "leal/docs"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("No se pudo cargar el archivo .env, usando variables del sistema: ", err)
		}
	}
}

// @title Tag Service API
// @version 1.0
// @description A Tag service API in Go using Gin Framework

// @host localhost:6060
// @BasePath /api/v1
func main() {
	r := gin.Default()

	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	api.InitServer(r)
}
