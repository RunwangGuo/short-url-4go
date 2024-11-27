package main

import "short-url-rw-github/src/controllers"

var serviceContainerObj *serviceContainer

type serviceContainer struct{}

type IServiceContainer interface {
	InjectLinkController() controllers.LinkController
}

func (sc *serviceContainer) InjectLinkController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyWriterController {
	// injecting service layer in controller
	return controllers.LinkController{}
}
