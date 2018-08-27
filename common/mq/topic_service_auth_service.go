package mq

import (
	"business/auth/cidl"

	"github.com/mz-eco/mz/kafka"
	"github.com/mz-eco/mz/log"
)

const (
	TOPIC_SERVICE_AUTH_SERVICE = "service-user-service"
)

const (
	IDENT_SERVICE_AUTH_SERVICE_INVALIDATE_TOKEN_CACHE = "invalidate_token_cache"
)

var (
	topicServiceAuthService *TopicServiceAuthService = nil
)

func GetTopicServiceAuthService() (topic *TopicServiceAuthService, err error) {
	if topicServiceAuthService != nil {
		topic = topicServiceAuthService
	}

	producer, err := kafka.NewAsyncProducer()
	if err != nil {
		log.Warnf("new async producer failed. %s", err)
		return
	}

	topicServiceAuthService = &TopicServiceAuthService{
		Producer: producer,
	}

	topic = topicServiceAuthService

	return
}

type TopicServiceAuthService struct {
	Producer *kafka.AsyncProducer
}

func (m *TopicServiceAuthService) send(ident string, msg interface{}) (err error) {
	err = m.Producer.SendMessage(TOPIC_SERVICE_AUTH_SERVICE, ident, msg)
	if err != nil {
		log.Warnf("send topic message failed. %s", err)
		return
	}
	return
}

type InvalidateTokenCacheMessage struct {
	AuthSiteType cidl.AuthSiteType // 站点 admin、city、wx_xcx
	Token        string
}

func (m *TopicServiceAuthService) InvalidateTokenCache(msg *InvalidateTokenCacheMessage) (err error) {
	return m.send(IDENT_SERVICE_AUTH_SERVICE_INVALIDATE_TOKEN_CACHE, msg)
}

func BroadcastInvalidateAuthAdminToken(token string) (err error) {

	// 广播token失效
	topicService, err := GetTopicServiceAuthService()
	if err != nil {
		log.Warnf("get topic service auth service failed. %s", err)
		return
	}

	err = topicService.InvalidateTokenCache(&InvalidateTokenCacheMessage{
		AuthSiteType: cidl.AuthSiteTypeAdmin,
		Token:        token,
	})
	if err != nil {
		log.Warnf("send topic service invalidate admin token cache failed. %s", err)
		return
	}

	return
}

func BroadcastInvalidateAuthOrgToken(token string) (err error) {

	// 广播token失效
	topicService, err := GetTopicServiceAuthService()
	if err != nil {
		log.Warnf("get topic service auth service failed. %s", err)
		return
	}

	err = topicService.InvalidateTokenCache(&InvalidateTokenCacheMessage{
		AuthSiteType: cidl.AuthSiteTypeOrg,
		Token:        token,
	})
	if err != nil {
		log.Warnf("send topic service invalidate admin token cache failed. %s", err)
		return
	}

	return
}

func BroadcastInvalidateAuthWxXcxToken(token string) (err error) {

	// 广播token失效
	topicService, err := GetTopicServiceAuthService()
	if err != nil {
		log.Warnf("get topic service auth service failed. %s", err)
		return
	}

	err = topicService.InvalidateTokenCache(&InvalidateTokenCacheMessage{
		AuthSiteType: cidl.AuthSiteTypeWxXcx,
		Token:        token,
	})
	if err != nil {
		log.Warnf("send topic service invalidate admin token cache failed. %s", err)
		return
	}

	return
}
