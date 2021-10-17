package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Nuttawut503/capstone-backend/db"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(viper.AllKeys())

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	if err := rdb.Set(ctx, "foo", "bar", 0).Err(); err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Key foo => ", val)

	_ = gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello World!")
	// })

	// r.POST("/auth", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "ping",
	// 	})
	// })
	// r.Run(":8080")
	generatedTime := time.Now()
	accessUUID := uuid.New().String()
	refreshUUID := uuid.New().String()
	userID := "1150"
	fmt.Println("access_uuid: ", accessUUID)
	fmt.Println("refresh_uuid: ", refreshUUID)
	AccessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"authorized":  true,
		"access_uuid": accessUUID,
		"user_id":     userID,
		"exp":         generatedTime.Add(time.Minute * 15).Unix(),
	}).SignedString([]byte(viper.GetString("ACCESS_SECRET")))
	fmt.Println("access: ", AccessToken)
	RefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"refresh_uuid": refreshUUID,
		"user_id":      userID,
		"exp":          generatedTime.Add(time.Hour * 6).Unix(),
	}).SignedString([]byte(viper.GetString("ACCESS_SECRET")))
	fmt.Println("refresh: ", RefreshToken)
	token, err := jwt.Parse(AccessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(viper.GetString("ACCESS_SECRET")), nil
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(token.Claims.(jwt.MapClaims))
	}
}
