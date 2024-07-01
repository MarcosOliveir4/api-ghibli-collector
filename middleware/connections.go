package middleware

import (
	"template-golang/database"

	"github.com/gin-gonic/gin"
)

func SetupConnections(db database.ConnectionsPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
	}
}
