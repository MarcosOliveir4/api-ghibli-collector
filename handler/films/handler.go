package films

import (
	"api-ghibli-collector/usecase/films"
	"net/http"

	"github.com/gin-gonic/gin"
)

func listFilms(c *gin.Context) {
	filmsList, err := films.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": filmsList,
	})
}
