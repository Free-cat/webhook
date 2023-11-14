package webhook

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
	def "webhook/internal/repository"
	repoModel "webhook/internal/repository/webhook/model"
)

var _ def.WebhookRepository = (*redisRepository)(nil)

type redisRepository struct {
	redisClient *redis.Client
	prefix      string
}

func NewRedisRepository(redisClient *redis.Client) *redisRepository {
	return &redisRepository{
		redisClient: redisClient,
		prefix:      "webhook_repo:",
	}
}

func (r *redisRepository) Create(ctx context.Context, uuid string, webhook *repoModel.Webhook) (*repoModel.Webhook, error) {
	record := &repoModel.Webhook{
		Uuid:      uuid,
		CreatedAt: time.Now(),
	}

	value, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}
	r.redisClient.Set(ctx, r.prefix+uuid, value, 0)

	return record, nil
}

func (r *redisRepository) Get(ctx context.Context, uuid string) (*repoModel.Webhook, error) {
	var webhook repoModel.Webhook
	value, err := r.redisClient.Get(ctx, r.prefix+uuid).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(value), &webhook)
	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

func (r *redisRepository) Delete(ctx context.Context, uuid string) error {
	return r.redisClient.Del(ctx, r.prefix+uuid).Err()
}
