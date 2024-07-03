package teste

import (
	"api-ghibli-collector/usecase/teste"
	"net/http"
	"strconv"

	domain "api-ghibli-collector/domain/teste"

	"github.com/gin-gonic/gin"
)

// @Summary List all testes
// @Description List all testes
// @Tags testes
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Teste
// @Router /testes [get]
func listTeste(c *gin.Context) {
	testes := teste.Listar(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "List of testes",
		"testes":  testes,
	})
}

// @Summary Create a teste
// @Description Create a teste
// @Tags testes
// @Accept json
// @Produce json
// @Param teste body domain.Teste true "Teste"
// @Success 201 {object} domain.Teste
// @Router /testes [post]
func createTeste(c *gin.Context) {
	var t domain.Teste
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	test, err := teste.Criar(c, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Teste created",
		"test":    test,
	})
}

// @Summary Find a teste by ID
// @Description Find a teste by ID
// @Tags testes
// @Accept json
// @Produce json
// @Param id path int true "Teste ID"
// @Success 200 {object} domain.Teste
// @Router /testes/{id} [get]
func findTesteByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	test, err := teste.BuscaPorID(c, idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Teste found",
		"test":    test,
	})
}
