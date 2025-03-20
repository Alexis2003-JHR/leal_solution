package api

import (
	"context"
	"fmt"
	"leal/cmd/db"
	"leal/internal/core/handlers"
	"leal/internal/core/middleware"
	"leal/internal/core/services"
	"leal/internal/repository"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	port = 6060
)

func InitServer(r *gin.Engine) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r.Use(func(c *gin.Context) {
		c.Set("context", ctx)
		c.Next()
	})

	db, err := db.Init()
	if err != nil {
		log.Fatal()
	}

	r.Use(middleware.ErrorHandler())

	repository := repository.NewPostgrestRepository(db)
	services := services.NewService(repository)
	handler := handlers.NewHandler(services)

	InitRouter(r, handler)

	r.Run(fmt.Sprintf(":%d", port))
}
