package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/industrix-todo/backend/internal/domain"
	"github.com/industrix-todo/backend/pkg/response"
)

type TodoHandler struct {
	service domain.TodoService
}

func NewTodoHandler(service domain.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/todos")
	{
		group.GET("", h.List)
		group.POST("", h.Create)
		group.GET("/:id", h.Get)
		group.PUT("/:id", h.Update)
		group.DELETE("/:id", h.Delete)
		group.PATCH("/:id/complete", h.ToggleComplete)
	}
}

func (h *TodoHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	catID, _ := strconv.ParseUint(c.Query("category_id"), 10, 32)

	filter := domain.TodoFilter{
		Page:       page,
		Limit:      limit,
		Search:     c.Query("search"),
		Status:     c.Query("status"),
		CategoryID: uint(catID),
		Priority:   c.Query("priority"),
		SortBy:     c.Query("sort_by"),
		SortOrder:  c.Query("sort_order"),
	}

	todos, total, err := h.service.ListTodos(c.Request.Context(), filter)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
    
    // Calculate total pages
    totalPages := int(math.Ceil(float64(total) / float64(limit)))
    if totalPages == 0 {
        totalPages = 1
    }

	pagination := response.Pagination{
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		TotalPages:  totalPages,
	}

	response.SuccessList(c, todos, pagination)
}

func (h *TodoHandler) Create(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_input", err.Error())
		return
	}

	if err := h.service.CreateTodo(c.Request.Context(), &todo); err != nil {
		response.Error(c, http.StatusBadRequest, "validation_failed", err.Error())
		return
	}

	response.Created(c, todo)
}

func (h *TodoHandler) Get(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_id", "invalid todo ID")
		return
	}

	todo, err := h.service.GetTodo(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "not_found", err.Error())
		return
	}

	response.Success(c, todo)
}

func (h *TodoHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_id", "invalid todo ID")
		return
	}

	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_input", err.Error())
		return
	}

	if err := h.service.UpdateTodo(c.Request.Context(), uint(id), &todo); err != nil {
		response.Error(c, http.StatusBadRequest, "validation_failed", err.Error())
		return
	}

	response.Success(c, todo)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_id", "invalid todo ID")
		return
	}

	if err := h.service.DeleteTodo(c.Request.Context(), uint(id)); err != nil {
		response.Error(c, http.StatusBadRequest, "validation_failed", err.Error())
		return
	}

	response.Success(c, gin.H{"deleted": true})
}

func (h *TodoHandler) ToggleComplete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid_id", "invalid todo ID")
		return
	}

	if err := h.service.ToggleComplete(c.Request.Context(), uint(id)); err != nil {
		response.Error(c, http.StatusBadRequest, "validation_failed", err.Error())
		return
	}

	response.Success(c, gin.H{"toggled": true})
}
