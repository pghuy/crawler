// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"context"
	"github.com/huypher/crawler/internal/infra"
)

// Injectors from wire.go:

func InitApplication(ctx context.Context) (*ApplicationContext, func(), error) {
	cronjob := infra.ProvideCronJob()
	config, err := infra.ProvideConfig()
	if err != nil {
		return nil, nil, err
	}
	cache, cleanup, err := infra.ProvideCache(config)
	if err != nil {
		return nil, nil, err
	}
	websocket := infra.ProvideWebsocketService()
	executor := infra.ProvideVozExecutor(cache, websocket)
	applicationContext := &ApplicationContext{
		cronjob:         cronjob,
		vozExec:         executor,
		websocketServer: websocket,
	}
	return applicationContext, func() {
		cleanup()
	}, nil
}
