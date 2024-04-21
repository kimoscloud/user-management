package middleware

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/configuration"
	"strings"
)

type TokenClaims struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func CheckInternalClient() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Kimos-Authorization")
		if tokenString == "" {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Unauthorized",
				},
			)
			context.Abort()
			return
		}

		authorizationHeaderSplitted := strings.Split(tokenString, "Bearer ")
		if len(authorizationHeaderSplitted) != 2 {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Invalid token",
				},
			)
			context.Abort()
			return
		}

		decodedToken, err := base64.StdEncoding.DecodeString(authorizationHeaderSplitted[1])
		if err != nil {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Invalid token",
				},
			)
			context.Abort()
			return
		}

		var tokenClaims TokenClaims
		err = json.Unmarshal(decodedToken, &tokenClaims)
		if err != nil {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Invalid token",
				},
			)
			context.Abort()
			return
		}

		clientConfig := configuration.GetInternalClientConfig()
		if tokenClaims.ClientId != clientConfig.GetInternalClientId() || tokenClaims.ClientSecret != clientConfig.GetInternalClientSecret() {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Unauthorized",
				},
			)
			context.Abort()
			return
		}

		context.Next()
	}
}
