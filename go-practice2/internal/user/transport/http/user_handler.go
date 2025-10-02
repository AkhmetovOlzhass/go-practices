package http

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "go-prcatice2/internal/user/service"
)

type UserHandler struct {
    svc *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
    return &UserHandler{svc: s}
}

func (h *UserHandler) Register(r *gin.Engine) {
    r.GET("/user", h.GetUser)
    r.POST("/user", h.CreateUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
    idStr := c.Query("id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    user, err := h.svc.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
        return
    }

    c.JSON(http.StatusOK, user)
}


func (h *UserHandler) CreateUser(c *gin.Context) {
    var input struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid name"})
        return
    }

    user, err := h.svc.CreateUser(input.Name)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}
