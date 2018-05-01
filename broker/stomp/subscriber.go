package stomp

import (
	"github.com/go-stomp/stomp"
	"github.com/micro/go-micro/broker"
)

type subscriber struct {
	opts  broker.SubscribeOptions
	topic string
	// Stomp subscription
	sub *stomp.Subscription
}

func (s *subscriber) Options() broker.SubscribeOptions {
	return s.opts
}

func (s *subscriber) Topic() string {
	return s.topic
}

func (s *subscriber) Unsubscribe() error {
	return s.sub.Unsubscribe()
}
