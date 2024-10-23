package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_users "ibercs/proto/users"
	"io"
	"io/ioutil"
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
	// Obtener el código de autorización y el code_verifier del frontend
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

	// Obtener client_id, client_secret y redirect_uri de variables de entorno
	clientID := os.Getenv("FACEIT_CLIENT_ID")
	clientSecret := os.Getenv("FACEIT_CLIENT_SECRET")
	redirectURI := os.Getenv("FACEIT_REDIRECT_URI")

	if clientID == "" || clientSecret == "" || redirectURI == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Configuración del servidor incompleta"})
		return
	}

	// Construir los datos para la solicitud de token, incluyendo el code_verifier
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("code_verifier", codeVerifier) // Asegúrate de incluir el code_verifier

	// Codificar las credenciales en Base64
	credentials := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	// Crear la solicitud HTTP POST al token endpoint de Faceit
	req, err := http.NewRequest("POST", "https://api.faceit.com/auth/v1/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud HTTP"})
		return
	}

	// Configurar los encabezados
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+credentials)

	// Realizar la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al realizar la solicitud HTTP"})
		return
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer la respuesta del servidor"})
		return
	}

	// Verificar el código de estado HTTP
	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Error al obtener tokens", "details": string(body)})
		return
	}

	// Decodificar la respuesta JSON
	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar la respuesta de tokens"})
		return
	}

	// Obtener información del usuario utilizando el access_token
	userInfo, err := getUserInfo(tokenResponse.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener información del usuario"})
		return
	}

	// Manejar la autenticación del usuario (crear sesión, etc.)

	// Como ejemplo, devolvemos la información del usuario
	c.JSON(http.StatusOK, gin.H{"user": userInfo})
}

func getUserInfo(accessToken string) (*UserInfo, error) {
	// Crear la solicitud HTTP GET al userinfo endpoint de Faceit
	req, err := http.NewRequest("GET", "https://api.faceit.com/auth/v1/resources/userinfo", nil)
	if err != nil {
		return nil, err
	}

	// Configurar los encabezados
	req.Header.Add("Authorization", "Bearer "+accessToken)

	// Realizar la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Verificar el código de estado HTTP
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error al obtener información del usuario: %s", string(body))
	}

	// Decodificar la respuesta JSON
	var userInfo UserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
