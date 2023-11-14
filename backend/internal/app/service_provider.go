package app

import (
	"github.com/redis/go-redis/v9"
	"webhook/internal/api/request"
	"webhook/internal/api/webhook"
	"webhook/internal/repository"
	requestRepository "webhook/internal/repository/request"
	webhookRepository "webhook/internal/repository/webhook"
	"webhook/internal/service"
	requestService "webhook/internal/service/request"
	webhookService "webhook/internal/service/webhook"
)

type serviceProvider struct {
	ginConfig *ginConfig

	redisClient *redis.Client

	webhookRepository    repository.WebhookRepository
	webhookService       service.WebhookService
	publicWebhookService service.WebhookPublicService
	webhookImpl          *webhook.Implementation

	requestRepository repository.RequestRepository
	requestService    service.RequestService
	requestImpl       *request.Implementation
}

type ginConfig struct {
	Host string
	Port string
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) init() {
	sp.initGinConfig()
	sp.initRedisClient()

	sp.initWebhookRepository()
	sp.initWebhookService()
	sp.initWebhookImpl()
	sp.initWebhookPublicService()

	sp.initRequestRepository()
	sp.initRequestService()
	sp.initRequestImpl()
}

func (sp *serviceProvider) initGinConfig() {
	sp.ginConfig = &ginConfig{
		Host: "localhost",
		Port: "8080",
	}
}

func (sp *serviceProvider) initRedisClient() {
	sp.redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func (sp *serviceProvider) initWebhookRepository() {
	sp.webhookRepository = webhookRepository.NewRedisRepository(sp.redisClient)
}

func (sp *serviceProvider) initWebhookService() {
	sp.webhookService = webhookService.NewService(sp.webhookRepository)
}

func (sp *serviceProvider) initWebhookPublicService() {
	sp.publicWebhookService = webhookService.NewPublicService(sp.webhookRepository)
}

func (sp *serviceProvider) initWebhookImpl() {
	sp.webhookImpl = webhook.NewImplementation(sp.webhookService)
}

func (sp *serviceProvider) getWebhookImpl() *webhook.Implementation {
	return sp.webhookImpl
}

func (sp *serviceProvider) initRequestRepository() {
	sp.requestRepository = requestRepository.NewRedisRepository(sp.redisClient)
}

func (sp *serviceProvider) initRequestService() {
	sp.requestService = requestService.NewService(sp.requestRepository)
}

func (sp *serviceProvider) initRequestImpl() {
	sp.requestImpl = request.NewImplementation(sp.requestService, sp.publicWebhookService)
}

func (sp *serviceProvider) getRequestImpl() *request.Implementation {
	return sp.requestImpl
}
