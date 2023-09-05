package routes

import (
	"echo-graphql-jwt/database"
	"echo-graphql-jwt/graph"
	"echo-graphql-jwt/handlers"
	"echo-graphql-jwt/utils"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("api/auth/register", handlers.Register)
	e.POST("api/auth/login", handlers.Login)

	e.POST("/query", graphqlHandler(), echojwt.WithConfig(utils.JwtConfig))
	e.GET("/", playgroundHandler())

	return e
}

func graphqlHandler() echo.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: database.DB}}))

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
