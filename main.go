package main

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Nuttawut503/capstone-backend/auth"
	"github.com/Nuttawut503/capstone-backend/db"
	"github.com/Nuttawut503/capstone-backend/graph"
	"github.com/Nuttawut503/capstone-backend/graph/generated"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

	r := gin.Default()
	r.GET("/", func() gin.HandlerFunc {
		return func(c *gin.Context) {
			playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
		}
	}())
	r.POST("/query", func(c *gin.Context) {
		handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})).ServeHTTP(c.Writer, c.Request)
	})
	auth.SetAuthRouter(r.Group("/auth"), rdb, ctx)
	authorized := r.Group("/secret", auth.GetMiddleware(rdb, ctx))
	authorized.GET("/me", func(c *gin.Context) {
		userID := c.MustGet("user_id").(string)
		c.JSON(200, gin.H{
			"response": "Welcome " + userID,
		})
	})
	r.Run(":8080")
}
