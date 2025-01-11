package publisher

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

type Publisher interface {
	Publish(ctx context.Context, topic, message string) error
}

type PubSubPublisher struct {
	Client *pubsub.Client
}

func NewPubSubPublisher(client *pubsub.Client) *PubSubPublisher {
	return &PubSubPublisher{Client: client}
}

func (p *PubSubPublisher) Publish(ctx context.Context, topic string, message string) error {

	t := p.Client.Topic(topic)

	if serverId, err := t.Publish(ctx, &pubsub.Message{Data: []byte(message)}).Get(ctx); err != nil {
		logrus.Errorln(err)
		return err
	} else {
		logrus.Info("servier id : ", serverId)
		return nil
	}

}
