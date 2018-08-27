package auth

import (
	"strings"

	"github.com/mz-eco/mz/http"
)

type AuthHandler struct {
	Urls          []string
	WhiteListUrls []string
	AccessControl func(ctx *http.Context) (canAccess bool)
}

// 判断是否可以继续访问
// wasChecked: 是否满足条件
// canContinue: 是否可以继续
func (m *AuthHandler) CanContinue(ctx *http.Context) (wasChecked bool, canContinue bool) {
	urlPath := ctx.Engine.Request.URL.Path
	canContinue = true
	wasChecked = true

	for _, url := range m.Urls {
		if strings.HasPrefix(urlPath, url) {
			// 检查是否在白名单中
			for _, url := range m.WhiteListUrls {
				if strings.HasPrefix(urlPath, url) {
					canContinue = true
					return
				}
			}

			if m.AccessControl(ctx) {
				canContinue = true
			} else {
				canContinue = false
			}
			return
		}
	}

	wasChecked = false
	canContinue = true
	return
}

// 满足一个条件即可，在最前的最先被验证
type ChainAuthHandler struct {
	AuthHandlers []*AuthHandler
}

func (m *ChainAuthHandler) AccessControlHandler(ctx *http.Context) (canContinue bool) {
	for _, authHandler := range m.AuthHandlers {
		wasChecked, authContinue := authHandler.CanContinue(ctx)
		if wasChecked { // 命中一个
			canContinue = authContinue
			return
		}
	}

	return
}
