package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GetMiddleware(rdb *redis.Client, ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var claims jwt.MapClaims
		if accessToken, err := extractToken(c.Request.Header); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		} else if t, err := verifyToken(accessToken, viper.GetString("ACCESS_SECRET")); err != nil || !isValid(t) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			claims = t.Claims.(jwt.MapClaims)
		}
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		userID, err := rdb.Get(ctx, accessUUID).Result()
		if err != nil || userID != claims["user_id"].(string) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}

func SetAuthRouter(r *gin.RouterGroup, rdb *redis.Client, ctx context.Context) {
	r.POST("/login", func(c *gin.Context) {
		loginForm := struct {
			UserID   string `json:"userid"`
			Password string `json:"password"`
		}{}
		if err := c.ShouldBindJSON(&loginForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "wrong format"})
			return
		}
		// +todo: check if loginForm.UserID exists in the database
		accessUUID := uuid.New().String()
		refreshUUID := uuid.New().String()
		now := time.Now()
		accessToken, err := createAccessToken(accessUUID, loginForm.UserID, now.Add(access_lifetime).Unix())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "to create access failed"})
			return
		}
		refreshToken, err := createRefreshToken(refreshUUID, loginForm.UserID, now.Add(refresh_lifetime).Unix())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "to create refresh failed"})
		}
		if err := rdb.Set(ctx, accessUUID, loginForm.UserID, access_lifetime).Err(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "to save access failed"})
			return
		}
		if err := rdb.Set(ctx, refreshUUID, loginForm.UserID, refresh_lifetime).Err(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "to save refresh failed"})
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"access_exp":    now.Add(access_lifetime).Unix(),
			"refresh_exp":   now.Add(refresh_lifetime).Unix(),
		})
	})
}
