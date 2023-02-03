package middleware

import (
	routing "github.com/gly-hub/fasthttp-routing"
	auth2 "github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/rpcserver"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/server/http"
	"strings"
)

func PermissionMiddleware() routing.Handler {
	return func(c *routing.Context) error {
		// 校验token
		var (
			req  = new(auth2.CheckTokenParams)
			resp = new(auth2.CheckTokenResp)
		)
		req.Token = strings.TrimPrefix(c.Header.Value("Authorization"), "Bearer ")
		err := application.RpcCall(c, rpcserver.ServerName, rpcserver.CheckToken, req, resp)
		if err != nil {
			c.Abort()
			return (&http.HttpController{}).Fail(c, err)
		}

		// 设置user_id TODO
		c.Header.Set("userId", resp.UserId)
		c.Header.Set("userName", resp.UserName)
		// TODO
		c.Header.SetInt("platform", 1)

		// 设置请求ip
		c.Header.Set("ip", c.RequestCtx.Conn().RemoteAddr().String())

		// 是否替换token
		if resp.NewToken != "" {
			c.Response.Header.Set("Authorization", "Bearer "+resp.NewToken)
		}

		return c.Next()
	}
}
