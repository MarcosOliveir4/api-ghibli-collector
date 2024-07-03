package films

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("/films", listFilms)
}
