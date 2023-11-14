package app

import (
	"context"
	"github.com/gin-gonic/gin"
	publicApi "webhook/api/public"
	api "webhook/api/webhook_v1"
)

type App struct {
	serviceProvider *serviceProvider
	ginServer       *gin.Engine
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.init(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (app *App) Run() error {
	return app.ginServer.Run(app.serviceProvider.ginConfig.Host + ":" + app.serviceProvider.ginConfig.Port)
}

func (app *App) init(ctx context.Context) error {
	inits := []func(context.Context) error{
		app.initServiceProvider,
		app.initGinServer,
		app.initRoutes,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initServiceProvider(_ctx context.Context) error {
	app.serviceProvider = newServiceProvider()
	app.serviceProvider.init()

	return nil
}

func (app *App) initGinServer(_ctx context.Context) error {
	app.ginServer = gin.Default()

	return nil
}

func (app *App) initRoutes(_ctx context.Context) error {
	apiGroup := app.ginServer.Group("/api/v1")
	inits := []func(apiGroup *gin.RouterGroup) error{
		app.initWebhookRoutes,
		app.initRequestRoutes,
	}
	publicGroup := app.ginServer.Group("/")
	publicInits := []func(apiGroup *gin.RouterGroup) error{
		app.initPublicRequestRoutes,
	}

	for _, f := range inits {
		err := f(apiGroup)
		if err != nil {
			return err
		}
	}

	for _, f := range publicInits {
		err := f(publicGroup)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initWebhookRoutes(apiGroup *gin.RouterGroup) error {
	webhookGroup := apiGroup.Group("/webhooks")
	api.RegisterWebhookRequests(webhookGroup, app.serviceProvider.getWebhookImpl())

	return nil
}

func (app *App) initRequestRoutes(apiGroup *gin.RouterGroup) error {
	requestGroup := apiGroup.Group("/requests")
	api.RegisterRequestRequests(requestGroup, app.serviceProvider.getRequestImpl())

	return nil
}

func (app *App) initPublicRequestRoutes(apiGroup *gin.RouterGroup) error {
	publicApi.RegisterRequestRequests(apiGroup, app.serviceProvider.getRequestImpl())

	return nil
}
