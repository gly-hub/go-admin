package middleware

import (
	routing "github.com/gly-hub/fasthttp-routing"
)

func PermissionMiddleware() routing.Handler {
	return func(c *routing.Context) error {
		// 校验token TODO

		// 设置user_id TODO
		c.Header.Set("user_name", "admin")
		c.Header.SetInt("platform", 1)

		return c.Next()
	}
}
