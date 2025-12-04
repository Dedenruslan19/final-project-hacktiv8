package middleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware validates JWT and stores claims into context
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	jwtSecretKey := os.Getenv("SECRET_KEY")
	config := echojwt.Config{
		SigningKey:    []byte(jwtSecretKey),
		NewClaimsFunc: func(c echo.Context) jwt.Claims { return jwt.MapClaims{} },
		ErrorHandler:  jwtErrorHandler,
	}

	return echojwt.WithConfig(config)(func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return jwtErrorHandler(c, echo.ErrUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return jwtErrorHandler(c, echo.ErrUnauthorized)
		}

		if id, ok := claims["id"].(float64); ok {
			c.Set("user_id", uint(id))
		}
		if role, ok := claims["role"].(string); ok {
			c.Set("role", role)
			c.Set("is_admin", role == "admin")
		} else {
			c.Set("is_admin", false)
		}
		if email, ok := claims["email"].(string); ok {
			c.Set("email", email)
		}

		return next(c)
	})
}

func jwtErrorHandler(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"message": "you are unauthorized",
		"status":  "error",
	})
}
