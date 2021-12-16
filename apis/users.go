package apis

import (
	"errors"
	"net/http"
	"strconv"
	"totality/users/client"
	"totality/users/totality"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

const (
	idParam = "id"
)

type Handler struct {
	client *client.Client
}

func (h *Handler) GetUserByID(c *gin.Context) {
	var param = c.Param(idParam)
	if param == "" {
		c.JSON(http.StatusBadRequest, errors.New("missing id"))

		return
	}

	var id, err = strconv.ParseInt(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid id"))

		return
	}
	var user totality.User
	user, err = h.client.GetUserByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsers(c *gin.Context) {
	var ids []int64
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var users, err = h.client.GetUsers(context.Background(), ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, users)
}

func NewHandler(c *client.Client) *Handler {
	return &Handler{
		client: c,
	}
}
