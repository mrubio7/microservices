package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_users "ibercs/proto/users"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

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

// TokenResponse representa la respuesta de Faceit al solicitar tokens
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

// UserInfo representa la información del usuario obtenida de Faceit
type UserInfo struct {
	GUID     string `json:"guid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Country  string `json:"country"`
	// Agrega otros campos según sea necesario
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

	code := jsonBody.Code
	codeVerifier := jsonBody.CodeVerifier

	if code == "" || codeVerifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Código de autorización o code_verifier faltante"})
		return
	}

	clientID := os.Getenv("FACEIT_CLIENT_ID")
	clientSecret := os.Getenv("FACEIT_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Configuración del servidor incompleta"})
		return
	}

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", "https://www.ibercs.com/my-profile")
	data.Set("code_verifier", codeVerifier)

	credentials := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	req, err := http.NewRequest("POST", "https://api.faceit.com/auth/v1/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud HTTP"})
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+credentials)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al realizar la solicitud HTTP"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer la respuesta del servidor"})
		return
	}

	logger.Info("Respuesta de FACEIT: %s", string(body))
	if resp.StatusCode != http.StatusOK {
		logger.Error("Error en la solicitud a FACEIT: %d, body: %s", resp.StatusCode, string(body))
		c.JSON(resp.StatusCode, gin.H{"error": "Error al obtener tokens", "details": string(body)})
		return
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar la respuesta de tokens"})
		return
	}

	userInfo, err := getUserInfo(tokenResponse.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener información del usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userInfo})
}

func getUserInfo(accessToken string) (*UserInfo, error) {
	req, err := http.NewRequest("GET", "https://api.faceit.com/auth/v1/resources/userinfo", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error al obtener información del usuario: %s", string(body))
	}

	var userInfo UserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
