package handlers

import (
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_users "ibercs/proto/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Users_Handlers struct {
	users_client pb_users.UserServiceClient
}

func NewUsersHandlers(usersClient pb_users.UserServiceClient) *Users_Handlers {
	return &Users_Handlers{
		users_client: usersClient,
	}
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

func (h *Users_Handlers) FaceitAuthCallback(c *gin.Context) {
	var jsonBody struct {
		Code         string `json:"code"`
		CodeVerifier string `json:"code_verifier"`
	}

	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de solicitud inválidos"})
		return
	}

	user, err := faceit.Auth(jsonBody.Code, jsonBody.CodeVerifier)
	if err != nil {
		logger.Error("Error authenticating user: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de solicitud inválidos"})
		return
	}

	var res *pb_users.User
	u, err := h.users_client.GetUserByFaceitId(c, &pb_users.GetUserRequest{Id: user.FaceitID})
	if u == nil {
		if err == gorm.ErrRecordNotFound {
			res, err = h.users_client.NewUser(c, &pb_users.NewUserRequest{FaceitId: user.FaceitID})
			if err != nil {
				c.JSON(http.StatusBadRequest, response.BuildError("Error creating user"))
				return
			}
		}

	}

	c.JSON(http.StatusOK, res)
}
