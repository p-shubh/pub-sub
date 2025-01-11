package config

import (
	"context"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func InitializePubSubClient(projectId, credentialsPath string) (*pubsub.Client, error) {

	//load the credentials from the path
	return pubsub.NewClient(context.Background(), projectId, option.WithAPIKey(credentialsPath))

}
