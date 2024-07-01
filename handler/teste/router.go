package teste

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("/teste", listTeste)
	r.GET("/teste/:id", findTesteByID)
	r.POST("/teste", createTeste)
}
