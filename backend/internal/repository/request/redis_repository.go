package request

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"sort"
	def "webhook/internal/repository"
	repoModel "webhook/internal/repository/request/model"
)

var _ def.RequestRepository = (*redisRepository)(nil)

type redisRepository struct {
	redisClient *redis.Client
	prefix      string
}

func NewRedisRepository(redisClient *redis.Client) *redisRepository {
	return &redisRepository{
		redisClient: redisClient,
		prefix:      "request_repo:",
	}
}

func (r *redisRepository) List(ctx context.Context, webhookUuid string) ([]*repoModel.Request, error) {
	values, err := r.redisClient.SMembers(ctx, r.prefix+webhookUuid).Result()
	if err != nil {
		return nil, err
	}
	requests := make([]*repoModel.Request, 0, len(values))
	for _, value := range values {
		if value == "" {
			continue
		}

		var request repoModel.Request
		err = json.Unmarshal([]byte(value), &request)
		if err != nil {
			log.Printf("error unmarshaling request: %v", err)
			continue
		}

		requests = append(requests, &request)
	}
	sort.Slice(requests[:], func(i, j int) bool {
		return requests[i].CreatedAt > requests[j].CreatedAt
	})

	return requests, nil
}

func (r *redisRepository) Create(ctx context.Context, webhookUuid string, request *repoModel.Request) (*repoModel.Request, error) {
	value, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r.redisClient.SAdd(ctx, r.prefix+webhookUuid, string(value))
	r.redisClient.Publish(ctx, r.prefix+webhookUuid, string(value))

	return request, nil
}

func (r *redisRepository) Get(ctx context.Context, webhookUuid string, uuid string) (*repoModel.Request, error) {
	requests, err := r.List(ctx, r.prefix+webhookUuid)
	if err != nil {
		return nil, err
	}
	for _, request := range requests {
		if request.Uuid == uuid {
			return request, nil
		}
	}
	return nil, nil
}

func (r *redisRepository) Delete(ctx context.Context, webhookUuid string, uuid string) error {
	return r.redisClient.SRem(ctx, r.prefix+webhookUuid, uuid).Err()
}

func (r *redisRepository) Subscribe(ctx context.Context, webhookUuid string) (<-chan *repoModel.Request, error) {
	ch := make(chan *repoModel.Request)
	redisChannel := r.redisClient.Subscribe(ctx, r.prefix+webhookUuid).Channel()

	go func() {
		for {
			select {
			case msg := <-redisChannel:
				var request repoModel.Request
				err := json.Unmarshal([]byte(msg.Payload), &request)
				if err != nil {
					log.Printf("error unmarshaling request: %v", err)
					continue
				}
				ch <- &request
			}
		}
	}()
	return ch, nil
}
