package auth

import (
	"business/auth/common/cache"

	"github.com/mz-eco/mz/http"
	"github.com/mz-eco/mz/log"
)

func NewWxXcxAuthHandler(urls []string, whiteListUrls []string) *AuthHandler {
	return &AuthHandler{
		Urls:          urls,
		WhiteListUrls: whiteListUrls,
		AccessControl: WxXcxAuthAccessControl,
	}
}

func WxXcxAuthAccessControl(ctx *http.Context) (canAccess bool) {
	canAccess = false
	uid := ctx.Session.Uid
	token := ctx.Session.Token

	if uid == "" || token == "" {
		canAccess = false
		log.Warnf("uid or token is empty")
		return
	}

	// TODO 假授权码
	if token == "123456" {
		canAccess = true
		return
	}

	_, err := cache.NewAuthCache().GetAuthWxXcx(uid, token)
	if err != nil {
		log.Warnf("get auth wx xcx failed. %s", err)
		return
	}

	canAccess = true

	return
}
