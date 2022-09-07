package middleware

import (
	"errors"
	"strings"

	"github.com/Nerzal/gocloak/v8"
	errorHandler "github.com/chocogem/bigquery-golang/pkg/api/error"
	"github.com/gin-gonic/gin"
)

var (
	clientId     = ""
	clientSecret = ""
	realm        = ""
	hostname     = ""
)

type AuthenticationHandler struct {
}

func NewAuthenticationHandler() *AuthenticationHandler {
	return &AuthenticationHandler{}
}

//Authen with keycloak
func (a *AuthenticationHandler) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := gocloak.NewClient(hostname)

		tokenWithBearer := c.GetHeader("Authorization")
		token := strings.ReplaceAll(tokenWithBearer, "Bearer ", "")
		if len(token) <= 0 {
			c.Error(errorHandler.NewErrorAuthentication(errors.New("Token not found in the request header")))
			c.Abort()
			return
		}

		authResult, err := client.RetrospectToken(c, token, clientId, clientSecret, realm)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		isActive := *authResult.Active
		if !isActive {
			c.Error(errorHandler.NewErrorAuthentication(errors.New("Token is not Active")))
			c.Abort()
			return
		}
		c.Next()

	}
}
