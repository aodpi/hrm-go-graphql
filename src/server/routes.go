package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aodpi/hrm-go-graphql/v2/controllers"
	"github.com/aodpi/hrm-go-graphql/v2/graph"
	"github.com/aodpi/hrm-go-graphql/v2/graph/generated"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://hrmdev.tk",
			"http://localhost:4200",
		},
		AllowMethods: []string{
			http.MethodPut,
			http.MethodPost,
			http.MethodGet,
			http.MethodDelete,
		},
		AllowHeaders:     []string{"Origin", "Authorization"},
		AllowCredentials: true,
	}))

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.GET("/graphql", playgroundHandler())
	router.POST("/graphql", graphqlHandler())

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/graphql")
	})

	return router
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
