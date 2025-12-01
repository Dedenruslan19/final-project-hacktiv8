package utils

import "github.com/labstack/echo/v4"

// GetUserID reads user id from context (set by auth middleware).
// Returns (0,false) when not present/invalid.
func GetUserID(c echo.Context) (uint, bool) {
	v := c.Get("user_id")
	switch t := v.(type) {
	case uint:
		return t, true
	case int:
		return uint(t), true
	case int64:
		return uint(t), true
	case float64:
		return uint(t), true
	default:
		return 0, false
	}
}

// IsAdmin reads "is_admin" flag from context (set by auth middleware).
func IsAdmin(c echo.Context) bool {
	v := c.Get("is_admin")
	if b, ok := v.(bool); ok && b {
		return true
	}
	return false
}
