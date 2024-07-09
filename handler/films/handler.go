package films

import (
	"api-ghibli-collector/usecase/films"
	"net/http"

	domain "api-ghibli-collector/domain/films"

	"github.com/gin-gonic/gin"
)

func listFilms(c *gin.Context) {
	var filter domain.FilmFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if filter.Total {
		total, err := films.ListTotal(c, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"total": total,
		})
		return
	}

	filmsList, err := films.List(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": filmsList,
	})
}
