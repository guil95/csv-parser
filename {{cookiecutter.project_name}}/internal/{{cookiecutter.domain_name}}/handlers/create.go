package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/{{cookiecutter.domain_name}}"
	"net/http"
)

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (p domainHandler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := p.uc.Create(c, {{cookiecutter.domain_name}}.{{cookiecutter.domain_struct}}{
		Name: req.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id.String()})
}
