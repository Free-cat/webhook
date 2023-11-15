package request

import (
	"context"
	"sync"
	def "webhook/internal/repository"
	repoModel "webhook/internal/repository/request/model"
)

var _ def.RequestRepository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	requests map[string]*requestArray
	m        sync.RWMutex
	Channels map[string]chan *repoModel.Request
}

type requestArray struct {
	Count    int
	Requests []*repoModel.Request
}

func NewInMemoryRequestRepository() *inMemoryRepository {
	return &inMemoryRepository{
		requests: make(map[string]*requestArray),
		Channels: make(map[string]chan *repoModel.Request),
	}
}

func (r *inMemoryRepository) List(_ context.Context, websocketUuid string) ([]*repoModel.Request, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	requests, ok := r.requests[websocketUuid]
	if !ok {
		return nil, nil
	}

	if requests.Count == 0 {
		return nil, nil
	}

	return requests.Requests, nil
}

func (r *inMemoryRepository) Create(_ context.Context, websocketUuid string, request *repoModel.Request) (*repoModel.Request, error) {
	r.m.Lock()
	defer r.m.Unlock()

	requests, ok := r.requests[websocketUuid]
	if !ok {
		requests = &requestArray{
			Count:    0,
			Requests: make([]*repoModel.Request, 0),
		}
	}

	requests.Count++
	requests.Requests = append(requests.Requests, request)

	if r.Channels != nil {
		ch, ok := r.Channels[websocketUuid]
		if ok {
			ch <- request
		}
	}

	r.requests[websocketUuid] = requests

	return request, nil
}

func (r *inMemoryRepository) Get(_ context.Context, websocketUuid string, uuid string) (*repoModel.Request, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	requests, ok := r.requests[websocketUuid]
	if !ok {
		return nil, nil
	}

	if requests.Count == 0 {
		return nil, nil
	}

	request := requests.Requests[requests.Count-1]

	return request, nil
}

func (r *inMemoryRepository) Delete(_ context.Context, websocketUuid string, uuid string) error {
	r.m.Lock()
	defer r.m.Unlock()

	requests, ok := r.requests[websocketUuid]
	if !ok {
		return nil
	}
	requests.Count--
	for i, request := range requests.Requests {
		if request.Uuid == uuid {
			requests.Requests = append(requests.Requests[:i], requests.Requests[i+1:]...)
			break
		}
	}

	return nil
}

func (r *inMemoryRepository) Subscribe(_ context.Context, websocketUuid string) (<-chan *repoModel.Request, error) {
	r.m.Lock()
	defer r.m.Unlock()

	ch := make(chan *repoModel.Request)
	r.Channels[websocketUuid] = ch

	return ch, nil
}
