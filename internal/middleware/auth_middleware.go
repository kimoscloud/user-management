package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kimoscloud/user-management-service/internal/core/auth"
	"strings"
)

// TODO add logger
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		appId := context.GetHeader("x-app-id")
		if tokenString == "" && appId == "" {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Unauthorized",
				},
			)
			context.Abort()
			return
		}

		if appId != "" {
			app, err := getApp(appId)
			if err != nil {
				context.AbortWithStatusJSON(
					401, gin.H{
						"message": "Invalid app ",
					},
				)
				context.Abort()
				return
			}
			context.Set("application", app)
		}

		if tokenString != "" {
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
			claims, err := auth.ValidateToken(
				authorizationHeaderSplitted[1],
			)
			if err != nil {
				context.AbortWithStatusJSON(
					401, gin.H{
						"message": "Unauthorized",
					},
				)
				context.Abort()
				return
			}
			context.Set("kimosUserId", claims.ID)
		}
		context.Next()
	}
}

func getApp(appId string) (string, error) {
	if appId == "organization" {
		return "kimos/org-app", nil
	}
	return "", errors.New("invalid app")
}
