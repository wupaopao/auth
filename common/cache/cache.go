package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"business/user/proxy/user"

	"github.com/mz-eco/mz/cache"
	"github.com/mz-eco/mz/errors"
	"github.com/mz-eco/mz/log"
)

type AuthCache struct {
	Cache *cache.Cache
}

const TTL_AUTH = time.Duration(1) * time.Hour

func NewAuthCache() *AuthCache {
	return &AuthCache{
		Cache: cache.NewRedis("auth"),
	}
}

func (m *AuthCache) KeyAuthAdmin(token string) string {
	return fmt.Sprintf("auth:tkn:admin:%s", token)
}

func (m *AuthCache) GetAuthAdmin(uid, token string) (auth *user.AuthAdmin, err error) {
	key := m.KeyAuthAdmin(token)
	strAuth, err := m.Cache.Get(key)
	if err != nil && err != cache.Nil {
		log.Warnf("get auth admin cache failed. %s", err)
		return

	} else if err == cache.Nil {
		userProxy := user.NewProxy("user-service")
		auth, err = userProxy.InnerUserAuthTokenInfoAdmin(&user.AskInnerUserAuthTokenInfoAdmin{
			UserID: uid,
			Token:  token,
		})
		if err != nil {
			log.Warnf("get auth admin from proxy failed. %s", err)
			return
		}

		byteAuth, errMarshal := json.Marshal(auth)
		if errMarshal != nil {
			err = errMarshal
			log.Warnf("marshal auth admin failed. %s", err)
			return
		}
		err = m.Cache.Set(key, byteAuth, TTL_AUTH)
		if err != nil {
			log.Warnf("set auth admin to cache failed. %s", err)
			return
		}

	} else {
		auth = &user.AuthAdmin{}
		err = json.Unmarshal([]byte(strAuth), auth)
		if err != nil {
			log.Warnf("unmarshal string auth admin failed. %s", err)
			return
		}
	}

	if uid != auth.UserId {
		err = errors.New("user id is not matched with auth.")
	}

	return
}

func (m *AuthCache) DeleteAuthAdmin(token string) (count int64, err error) {
	return m.Cache.Delete(m.KeyAuthAdmin(token))
}

func (m *AuthCache) KeyAuthCity(token string) string {
	return fmt.Sprintf("auth:tkn:city:%s", token)
}

func (m *AuthCache) GetAuthCity(uid, token string) (auth *user.AuthCity, err error) {
	key := m.KeyAuthCity(token)
	strAuth, err := m.Cache.Get(key)
	if err != nil && err != cache.Nil {
		log.Warnf("get auth admin cache failed. %s", err)
		return

	} else if err == cache.Nil { // 没有记录，则从user服务中获取授权信息并缓存到本地
		userProxy := user.NewProxy("user-service")
		auth, err = userProxy.InnerUserAuthTokenInfoOrg(&user.AskInnerUserAuthTokenInfoOrg{
			UserID: uid,
			Token:  token,
		})
		if err != nil {
			log.Warnf("get auth admin from proxy failed. %s", err)
			return
		}

		byteAuth, errMarshal := json.Marshal(auth)
		if errMarshal != nil {
			err = errMarshal
			log.Warnf("marshal auth admin failed. %s", err)
			return
		}
		err = m.Cache.Set(key, byteAuth, TTL_AUTH)
		if err != nil {
			log.Warnf("set auth city failed. %s", err)
			return
		}

	} else {
		auth = &user.AuthCity{}
		err = json.Unmarshal([]byte(strAuth), auth)
		if err != nil {
			log.Warnf("unmarshal string auth city failed. %s", err)
			return
		}
	}

	// 检查token与uid是否匹配
	if uid != auth.UserId {
		err = errors.New("user id is not matched with auth")
	}

	return
}

func (m *AuthCache) DeleteAuthCity(token string) (count int64, err error) {
	return m.Cache.Delete(m.KeyAuthCity(token))
}

func (m *AuthCache) KeyAuthWxXcx(token string) string {
	return fmt.Sprintf("auth:tkn:wx_xcx:%s", token)
}

func (m *AuthCache) GetAuthWxXcx(uid, token string) (auth *user.AuthWxXcx, err error) {
	key := m.KeyAuthWxXcx(token)
	strAuth, err := m.Cache.Get(key)
	if err != nil && err != cache.Nil {
		log.Warnf("get auth admin cache failed. %s", err)
		return
	} else if err == cache.Nil {
		userProxy := user.NewProxy("user-service")
		auth, err = userProxy.InnerUserAuthTokenInfoWxXcx(&user.AskInnerUserAuthTokenInfoWxXcx{
			UserID: uid,
			Token:  token,
		})
		byteAuth, errMarshal := json.Marshal(auth)
		if errMarshal != nil {
			err = errMarshal
			log.Warnf("marshal auth wx_xcx failed. %s", err)
			return
		}
		err = m.Cache.Set(key, byteAuth, TTL_AUTH)
		if err != nil {
			log.Warnf("set auth wx_xcx failed. %s", err)
			return
		}

	} else {
		auth = &user.AuthWxXcx{}
		err = json.Unmarshal([]byte(strAuth), auth)
		if err != nil {
			log.Warnf("unmarshal string auth wx_xcx failed. %s", err)
			return
		}
	}

	if uid != auth.UserId {
		err = errors.New("user id is not matched with auth")
	}

	return
}

func (m *AuthCache) DeleteAuthWxXcx(token string) (count int64, err error) {
	return m.Cache.Delete(m.KeyAuthWxXcx(token))
}
