package auth

import (
	"business/auth/common/cache"

	"github.com/mz-eco/mz/http"
	"github.com/mz-eco/mz/log"
)

func NewOrgAuthHandler(urls []string, whiteListUrls []string) *AuthHandler {
	return &AuthHandler{
		Urls:          urls,
		WhiteListUrls: whiteListUrls,
		AccessControl: OrgAuthAccessControl,
	}
}

func OrgAuthAccessControl(ctx *http.Context) (canAccess bool) {
	canAccess = false
	uid := ctx.Session.Uid
	token := ctx.Session.Token

	//for test
	log.Infof("uid:%s,token:%s",uid,token)
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

	log.Infof("auth token 1")
	_, err := cache.NewAuthCache().GetAuthCity(uid, token)
	if err != nil {
		log.Warnf("get auth city failed. %s", err)
		return
	}

	canAccess = true

	return
}
