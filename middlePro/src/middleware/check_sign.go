package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/labstack/echo/v4"
	"github.com/pengk/summer/config"
	"github.com/pengk/summer/util/constant"
	"github.com/pengk/summer/util/json"
	"github.com/pengk/summer/util/sign"
)

func CheckApiSign() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			params, err := ioutil.ReadAll(ctx.Request().Body)
			if err != nil {
				return constant.SignatureErr
			}
			m := make(map[string]interface{})
			err = json.Cjson.Unmarshal(params, &m)
			signature, ok := m["signature"]
			if !ok {
				return constant.SignatureErr
			}
			checkSignature, err := sign.Get(m, config.GetApiAuthToken())

			if err != nil {
				return constant.SignatureErr
			}
			if checkSignature != signature {
				return constant.SignatureErr
			}
			ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(params))
			return next(ctx)
		}
	}
}
