package handlers

import (
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	"ibercs/pkg/twitch"
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

func NewUsersHandlers(usersClient pb_users.UserServiceClient) *Users_Handlers {
	return &Users_Handlers{
		users_client: usersClient,
	}
}

func (h *Users_Handlers) GetUser(c *gin.Context) {
	id := c.Query("id")
	nickname := c.Query("nickname")

	if id == "" && nickname == "" {
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid params"))
		return
	}

	var res *pb_users.User
	var err error
	if id != "" {
		res, err = h.users_client.GetUser(c, &pb_users.GetUserRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting user"))
			return
		}
	} else {
		res, err = h.users_client.GetUserByPlayerNickname(c, &pb_users.GetUserRequest{Id: nickname})
		if err != nil {
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting user"))
			return
		}
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (h *Users_Handlers) GetStreams(c *gin.Context) {
	streams, err := h.users_client.GetAllStreams(c, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting streams"))
		return
	}

	var res []*twitch.Channel
	for _, channel := range streams.Streams {
		ch := twitch.GetStreamData(channel.Stream, channel.Name)
		if ch != nil {
			res = append(res, ch)
		}
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
	res, err = h.users_client.GetUserByFaceitId(c, &pb_users.GetUserRequest{Id: user.FaceitID})
	if res == nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			res, err = h.users_client.NewUser(c, &pb_users.NewUserRequest{FaceitId: user.FaceitID})
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
		Session: session.Response,
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", responseData))
}

func (h *Users_Handlers) UpdateProfile(c *gin.Context) {
	identity, exist := c.Get("identity")
	if !exist {
		c.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
		return
	}

	var req struct {
		Twitter string `json:"twitter"`
		Twitch  string `json:"twitch"`
		Desc    string `json:"desc"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid data"))
		return
	}

	id := strconv.Itoa(identity.(int))
	user, err := h.users_client.GetUser(c, &pb_users.GetUserRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BuildError("Internal error"))
		return
	}

	user.Description = req.Desc
	user.Twitch = req.Twitch
	user.Twitter = req.Twitter

	res, err := h.users_client.UpdateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BuildError("Internal error"))
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
