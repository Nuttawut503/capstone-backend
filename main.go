package main

import (
	_ "context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/Nuttawut503/capstone-backend/auth"
	"github.com/Nuttawut503/capstone-backend/db"
	"github.com/Nuttawut503/capstone-backend/graph"
	"github.com/Nuttawut503/capstone-backend/graph/generated"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     viper.GetString("REDIS_URL"),
	// 	Password: "",
	// 	DB:       0,
	// })
	// ctx := context.Background()

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func() gin.HandlerFunc {
		return func(c *gin.Context) {
			playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
		}
	}())
	//auth.SetAuthRouter(r.Group("/auth"), rdb, ctx)
	r.POST("/query" /**auth.GetMiddleware(rdb, ctx),*/, func(c *gin.Context) {
		handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Client: client}})).ServeHTTP(c.Writer, c.Request)
	})
	r.Run(":8080")
}
