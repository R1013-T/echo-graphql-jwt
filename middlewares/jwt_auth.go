package middlewares

import (
	"bytes"
	"echo-graphql-jwt/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"strings"
)

func SkipAuthForLoginAndRegister(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			body, _ := ioutil.ReadAll(c.Request().Body)
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
			// Check if the request body contains the login or register mutations
			if strings.Contains(string(body), "mutation Register") || strings.Contains(string(body), "mutation Login") {
				// If it's a login or register mutation, call the next handler directly
				return next(c)
			}
		}
		// Otherwise, apply the JWT middleware
		jwtMiddleware := echojwt.WithConfig(utils.JwtConfig)
		return jwtMiddleware(next)(c)
	}
}
