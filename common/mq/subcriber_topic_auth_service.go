package mq

import (
	"encoding/json"

	"business/auth/cidl"
	"business/auth/common/cache"

	"github.com/mz-eco/mz/kafka"
	"github.com/mz-eco/mz/log"
	"github.com/mz-eco/mz/settings"
)

var (
	topicAuthServiceGroupSetting kafka.TopicGroupSetting
)

func init() {
	settings.RegisterWith(func(viper *settings.Viper) error {
		err := viper.Unmarshal(&topicAuthServiceGroupSetting)
		if err != nil {
			panic(err)
			return err
		}
		return nil
	}, "kafka.subscribe.service_auth_service")
}

type TopicAuthServiceHandler struct {
	kafka.TopicHandler
}

func NewTopicAuthServiceHandler() (handler *TopicAuthServiceHandler, err error) {
	handler = &TopicAuthServiceHandler{
		TopicHandler: kafka.TopicHandler{
			Topics:  []string{TOPIC_SERVICE_AUTH_SERVICE},
			Brokers: topicAuthServiceGroupSetting.Address,
			Group:   topicAuthServiceGroupSetting.Group,
		},
	}

	handler.MessageHandle = handler.handleMessage

	return
}

func (m *TopicAuthServiceHandler) handleMessage(identMessage *kafka.IdentMessage) (err error) {
	switch identMessage.Ident {
	case IDENT_SERVICE_AUTH_SERVICE_INVALIDATE_TOKEN_CACHE:
		invalidateMessage := &InvalidateTokenCacheMessage{}
		err = json.Unmarshal(identMessage.Msg, invalidateMessage)
		if err != nil {
			log.Warnf("unmarshal invalidate token cache message failed. %s", err)
			return
		}

		err = m.InvalidateTokenCache(invalidateMessage)
		if err != nil {
			log.Warnf("invalidate token cache failed. %s", err)
			return
		}
	}

	return
}

func (m *TopicAuthServiceHandler) InvalidateTokenCache(msg *InvalidateTokenCacheMessage) (err error) {
	authCache := cache.NewAuthCache()
	switch msg.AuthSiteType {
	case cidl.AuthSiteTypeAdmin:
		_, err = authCache.DeleteAuthAdmin(msg.Token)
	case cidl.AuthSiteTypeOrg:
		_, err = authCache.DeleteAuthCity(msg.Token)
	case cidl.AuthSiteTypeWxXcx:
		_, err = authCache.DeleteAuthWxXcx(msg.Token)
	}
	return
}
