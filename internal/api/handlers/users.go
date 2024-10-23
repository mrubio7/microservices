package handlers

import (
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_users "ibercs/proto/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users_Handlers struct {
	users_client pb_users.UserServiceClient
}

func NewUsersHandlers(usersClient pb_users.UserServiceClient) *Users_Handlers {
	return &Users_Handlers{
		users_client: usersClient,
	}
}

func (h *Users_Handlers) NewUser(c *gin.Context) {
	var input struct {
		FaceitId string `json:"faceit_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error creating user"))
		return
	}

	res, err := h.users_client.NewUser(c, &pb_users.NewUserRequest{FaceitId: input.FaceitId})
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError("Error creating user"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (h *Users_Handlers) GetUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid ID"))
		return
	}

	res, err := h.users_client.GetUser(c, &pb_users.GetUserRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting user"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}
