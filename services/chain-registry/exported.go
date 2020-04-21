package chainregistry

import (
	"context"
	"sync"

	"github.com/containous/traefik/v2/pkg/log"
	"github.com/spf13/viper"
)

var initOnce = &sync.Once{}
var server *EnvelopStoreServer

func Init(ctx context.Context) {
	cfg := NewConfigFromViper(viper.GetViper())
	initOnce.Do(func() {
		var err error
		server, err = NewServer(ctx, &cfg)
		if err != nil {
			log.FromContext(ctx).WithError(err).Fatalf("Could not create envelope store application")
		}
	})
}

func Start(ctx context.Context) error {
	Init(ctx)
	return server.Start()
}

func Stop(ctx context.Context) error {
	Init(ctx)
	return server.Stop()
}