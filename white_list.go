package auth

import (
	"github.com/mz-eco/mz/http"
)

func NewWhiteListAuthHandler(urls []string) *AuthHandler {
	return &AuthHandler{
		Urls:          urls,
		AccessControl: WhiteListAuthAccessControl,
	}
}

func WhiteListAuthAccessControl(ctx *http.Context) (canAccess bool) {
	canAccess = true
	return
}
