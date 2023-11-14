package webhook

import (
	"context"
	"fmt"
	"sync"
	"time"
	def "webhook/internal/repository"
	repoModel "webhook/internal/repository/webhook/model"
)

var _ def.WebhookRepository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	data map[string]*repoModel.Webhook
	m    sync.RWMutex
}

func NewInMemoryWebhookRepository() *inMemoryRepository {
	return &inMemoryRepository{
		data: make(map[string]*repoModel.Webhook),
	}
}

func (r *inMemoryRepository) Create(_ context.Context, uuid string, webhook *repoModel.Webhook) (*repoModel.Webhook, error) {
	r.m.Lock()
	defer r.m.Unlock()

	record := &repoModel.Webhook{
		Uuid:      uuid,
		CreatedAt: time.Now(),
	}
	r.data[uuid] = record

	return record, nil
}

func (r *inMemoryRepository) Get(_ context.Context, uuid string) (*repoModel.Webhook, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	webhook, ok := r.data[uuid]
	if !ok {
		return nil, fmt.Errorf("webhook not found")
	}

	return webhook, nil
}

func (r *inMemoryRepository) Delete(_ context.Context, uuid string) error {
	r.m.Lock()
	defer r.m.Unlock()

	delete(r.data, uuid)

	return nil
}
