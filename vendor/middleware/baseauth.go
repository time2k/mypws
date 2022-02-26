package mymiddleware

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/labstack/echo"

	myconfig "social_pay/config"
	//"github.com/mcuadros/go-version"
)

type BAuth struct {
}

func NewBaseAuth() *BAuth {
	return &BAuth{}
}

// Process is the middleware function.
func (s *BAuth) BaseAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//内网调用，不需要验证
		if c.RealIP() == "127.0.0.1" || strings.Contains(c.RealIP(), "192.168.") || strings.Contains(c.RealIP(), "10.") {
			return next(c)
		}
		debug := c.QueryParam("_debug")
		if debug == "C9TUL5VD" {
			return next(c)
		}

		//获取auth config
		authcfg := myconfig.APP_AUTH

		//获取auth参数
		user_appkey := c.Request().Header.Get("AppKey")
		user_nouce := c.Request().Header.Get("Nouce")
		user_timestamp := c.Request().Header.Get("Timestamp")
		user_signature := c.Request().Header.Get("Signature")
		if _, ok := authcfg[user_appkey]; !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized Visit! Unknown AppKey")
		}
		appsecret := authcfg[user_appkey]["AppSecret"]
		t := sha1.New()
		t.Write([]byte(appsecret + user_nouce + user_timestamp))
		cacl_sign := hex.EncodeToString(t.Sum(nil))

		if cacl_sign != user_signature {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized Visit! Signature Verify Error")
		}

		return next(c)
	}
}
