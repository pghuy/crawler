// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"context"
	"github.com/huypher/crawler/internal/infra"
)

// Injectors from wire.go:

func InitApplication(ctx context.Context) (*ApplicationContext, func(), error) {
	config, err := infra.ProvideConfig()
	if err != nil {
		return nil, nil, err
	}
	rabbitmqConsumer := infra.ProvideRabbitmqConsumer(config)
	consumer := infra.ProvideConsumer(config, rabbitmqConsumer)
	applicationContext := &ApplicationContext{
		consumer: consumer,
	}
	return applicationContext, func() {
	}, nil
}
