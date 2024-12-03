package main

import (
	"short-url-4go/src/config"
	"short-url-4go/src/controllers"
	"short-url-4go/src/interfaces"
	"short-url-4go/src/services"
	"sync"
)

var (
	serviceContainerObj *serviceContainer
	containerOnce       sync.Once
)

type serviceContainer struct{}

type IServiceContainer interface {
	InjectLinkController(dbClient interfaces.IDataAccessLayer, cache interfaces.ICacheLayer) controllers.LinkController
}

func (sc *serviceContainer) InjectLinkController(dbClient interfaces.IDataAccessLayer, cache interfaces.ICacheLayer) controllers.LinkController {
	// injecting service layer in controller
	return controllers.LinkController{
		ILinkService: &services.LinkService{
			IDataAccessLayer: dbClient,         //injecting db client
			ICacheLayer:      cache,            //injecting redisClient
			Logger:           config.ZapLogger, //injecting zapLogger
		},
		Logger: config.ZapLogger,
	}
}

func ServiceContainer() IServiceContainer {
	if serviceContainerObj == nil {
		containerOnce.Do(func() {
			serviceContainerObj = &serviceContainer{}
		})
	}
	return serviceContainerObj
}
