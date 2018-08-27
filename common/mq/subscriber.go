package mq

import (
	"github.com/mz-eco/mz/kafka"
	"github.com/mz-eco/mz/log"
)

func NewSubscriber() (subscriber *kafka.Subscriber, err error) {
	subscriber = &kafka.Subscriber{}
	topicAuthHandler, err := NewTopicAuthServiceHandler()
	if err != nil {
		log.Warnf("new topic auth service handler failed. %s", err)
		return
	}

	subscriber.TopicHandlers = append(subscriber.TopicHandlers, topicAuthHandler)
	return
}
