package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/industrix-todo/backend/internal/domain"
	"github.com/industrix-todo/backend/pkg/response"
)

type CategoryHandler struct {
	service domain.CategoryService
}

func NewCategoryHandler(service domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/categories")
	{
		group.GET("", h.List)
		group.POST("", h.Create)
		group.PUT("/:id", h.Update)
		group.DELETE("/:id", h.Delete)
	}
}

func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.service.ListCategories(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	response.Success(c, categories)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_input", err.Error())
		return
	}

	if err := h.service.CreateCategory(c.Request.Context(), &category); err != nil {
		response.Error(c, http.StatusBadRequest, "validation_failed", err.Error())
		return
	}

	response.Created(c, category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_id", "invalid category ID")
		return
	}

	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_input", err.Error())
		return
	}

	if err := h.service.UpdateCategory(c.Request.Context(), uint(id), &category); err != nil {
		response.Error(c, http.StatusBadRequest, "validation_failed", err.Error())
		return
	}

	response.Success(c, category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_id", "invalid category ID")
		return
	}

	if err := h.service.DeleteCategory(c.Request.Context(), uint(id)); err != nil {
		response.Error(c, http.StatusBadRequest, "validation_failed", err.Error())
		return
	}

	response.Success(c, gin.H{"deleted": true})
}
