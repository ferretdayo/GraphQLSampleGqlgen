package api

import "github.com/gin-gonic/gin"

const (
	userContextKey  = "user_context_key"
	tokenContextKey = "token_context_key"
)

func SetUserID(context *gin.Context, userID uint) {
	context.Set(userContextKey, userID)
}

func GetUserID(context *gin.Context) uint {
	return context.MustGet(userContextKey).(uint)
}

func SetTokenID(context *gin.Context, tokenID string) {
	context.Set(tokenContextKey, tokenID)
}

func GetTokenID(context *gin.Context) string {
	return context.MustGet(tokenContextKey).(string)
}
