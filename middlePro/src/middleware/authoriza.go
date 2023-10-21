package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Resp struct{}
type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	RequestID  string      `json:"request_id"`
}

func AuthorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("授权错误")
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization header missing"})
		}
		// 在此处添加你的授权逻辑，例如检查令牌的有效性

		// 如果授权通过，可以将用户信息存储在上下文中，以供后续处理程序使用
		// c.Set("user", user)
		fmt.Println("在此处添加你的授权逻辑")
		return next(c)
	}
}
