package routes

import (
	"context"
	"echo-graphql-jwt/database"
	"echo-graphql-jwt/graph"
	"echo-graphql-jwt/middlewares"
	"echo-graphql-jwt/utils"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			gqlCtx := context.WithValue(c.Request().Context(), utils.EchoContextKey, c)
			c.SetRequest(c.Request().WithContext(gqlCtx))
			return next(c)
		}
	})

	e.POST("/query", graphqlHandler(), middlewares.SkipAuthForLoginAndRegister)
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
