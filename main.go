package main

import (
	"api-ghibli-collector/database"
	"api-ghibli-collector/docs"
	"api-ghibli-collector/handler/films"
	"api-ghibli-collector/handler/teste"
	"api-ghibli-collector/middleware"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	dbInstance := &database.DB{}
	err := dbInstance.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	pools, err := dbInstance.CreatePool()
	if err != nil {
		log.Fatal(err)
	}

	dbInstance.CreateTables(pools, context.Background())
	defer dbInstance.CloseConnection()

	r := gin.Default()
	v1 := r.Group("v1")

	v1.Use(
		middleware.SetupConnections(pools),
	)

	teste.Router(v1)
	films.Router(v1)
	setupSwagDocs()
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on port 8080")
}

func setupSwagDocs() {
	docs.SwaggerInfo.Title = "Template Golang API"
	docs.SwaggerInfo.Description = "This is template golang app server"
	docs.SwaggerInfo.Version = "1.0"

	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
}
