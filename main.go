package main

import (
	"fmt"

	"github.com/MarcosOliveir4/api-ghibli-collector/db"
	"github.com/MarcosOliveir4/api-ghibli-collector/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := db.ConnectDB()
	defer db.Close()

	handlers.SetupRoutes(r, db)

	fmt.Println("Conexão com o banco de dados SQLite estabelecida com sucesso!")
	r.Run(":8080")
}
