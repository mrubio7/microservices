package handlers

import (
	"ibercs/cmd/api_gateway/api_app/requests"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_users "ibercs/proto/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Users_Handlers struct {
	users_client pb_users.UserServiceClient
}

func NewUsersHandlers(client pb_users.UserServiceClient) *Users_Handlers {
	return &Users_Handlers{
		users_client: client,
	}
}

func (h *Users_Handlers) Get(c *gin.Context) {
	idStr := c.Query("id")
	faceitId := c.Query("faceit_id")

	if idStr == "" && faceitId == "" {
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid params"))
		return
	}

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.BuildError("Invalid id"))
			return
		}

		res, err := h.users_client.GetUserById(c, &pb_users.GetUserByIdRequest{Id: int32(id)})
		if err != nil {
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting user by id"))
			return
		}
		c.JSON(http.StatusOK, response.BuildOk("Ok", res))
		return
	}

	if faceitId != "" {
		res, err := h.users_client.GetUserByFaceitId(c, &pb_users.GetUserRequest{Id: faceitId})
		if err != nil {
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting user by faceit_id"))
			return
		}
		c.JSON(http.StatusOK, response.BuildOk("Ok", res))
		return
	}

	c.JSON(http.StatusBadRequest, response.BuildError("Invalid params"))
}

func (h *Users_Handlers) Update(c *gin.Context) {
	identity, exist := c.Get("identity")
	if !exist {
		c.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
		return
	}

	var req requests.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid request"))
		return
	}

	userToUpdate, err := h.users_client.GetUserById(c, &pb_users.GetUserByIdRequest{Id: int32(identity.(int))})
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting user"))
		return
	}

	pbUser, err := req.ToProto(userToUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid request"))
		return
	}

	res, err := h.users_client.UpdateUser(c, pbUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BuildError("Error updating user"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

// Access
func (h *Users_Handlers) Login(c *gin.Context) {
	identity, identityExist := c.Get("identity")
	_, tokenExist := c.Get("token")

	if !identityExist || !tokenExist {
		c.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
		return
	}

	res, err := h.users_client.GetUserById(c, &pb_users.GetUserByIdRequest{Id: int32(identity.(int))})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting user"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (h *Users_Handlers) Logout(c *gin.Context) {
	identity, exist := c.Get("identity")
	if !exist {
		c.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
		return
	}

	_, err := h.users_client.DeleteSession(c, &pb_users.NewSessionRequest{Id: int32(identity.(int))})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BuildError("Internal error"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", nil))
}

// Streams
func (h *Users_Handlers) GetStreams(c *gin.Context) {
	streams, err := h.users_client.GetAllStreams(c, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting streams"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", streams))
}

// Auth
func (h *Users_Handlers) AuthCallback_Faceit(c *gin.Context) {
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
	res, err = h.users_client.GetUserByFaceitId(c, &pb_users.GetUserRequest{Id: user.FaceitId})
	if res == nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			res, err = h.users_client.NewUser(c, &pb_users.NewUserRequest{FaceitId: user.FaceitId})
			if err != nil {
				c.JSON(http.StatusBadRequest, response.BuildError("Error creating user"))
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, response.BuildError("Unexpected error"))
			return
		}
	}

	h.users_client.DeleteSession(c, &pb_users.NewSessionRequest{Id: res.ID})
	session, err := h.users_client.NewSession(c, &pb_users.NewSessionRequest{Id: res.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BuildError("Unexpected error"))
		return
	}

	responseData := struct {
		User    *pb_users.User
		Session string
	}{
		User:    res,
		Session: session.Token,
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", responseData))
}
