package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	v1 := r.Group("/v1")
	{
		films(v1, db)
	}
}

func films(r *gin.RouterGroup, db *sql.DB) {
	r.GET("/films", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
