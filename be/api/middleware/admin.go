package middleware

import (
	"milestone3/be/internal/utils"

	"github.com/labstack/echo/v4"
)

// RequireAdmin ensures request has is_admin=true in context (set by auth middleware)
func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !utils.IsAdmin(c) {
			return utils.ForbiddenResponse(c, "forbidden")
		}
		return next(c)
	}
}
