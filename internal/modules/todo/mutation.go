package todo

import (
	"errors"
	"log/slog"
	"net/http"

	"gin-backend/internal/domain"
	"gin-backend/internal/domain/todo/usecase"
	"github.com/gin-gonic/gin"
)

type Mutation struct {
	createTodoUC *usecase.CreateTodoUseCase
}

func NewMutation(c *usecase.CreateTodoUseCase) *Mutation {
	return &Mutation{createTodoUC: c}
}

func (m *Mutation) Create(c *gin.Context) {
	var input usecase.CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := m.createTodoUC.Execute(c.Request.Context(), input)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func handleError(c *gin.Context, err error) {
	if errors.Is(err, domain.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, domain.ErrConflict) {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	slog.Error("unexpected error", "error", err, "path", c.Request.URL.Path)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
