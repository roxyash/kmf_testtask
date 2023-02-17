package repository

import (
	"github.com/roxyash/kmf_testtask/proxy_service/internal/repository/model"
	"sync"
)

type responseMap struct {
	mu sync.RWMutex
	m  map[int]model.ProxyModel
}

func NewResponseMap(mu sync.RWMutex, m map[int]model.ProxyModel) *responseMap {
	return &responseMap{
		mu: mu,
		m:  m,
	}
}

type StorageRepo struct {
	mapA *responseMap
}

func NewSomeStorageRepo() *StorageRepo {
	// Initialize map
	m := make(map[int]model.ProxyModel)
	mapA := NewResponseMap(sync.RWMutex{}, m)
	return &StorageRepo{
		mapA: mapA,
	}
}

func (s *StorageRepo) SetProxyData(response model.ProxyModel) model.ProxyModel {
	s.mapA.mu.Lock()
	id := len(s.mapA.m) + 1
	proxyModel := model.ProxyModel{
		ID:      id,
		Status:  response.Status,
		Headers: response.Headers,
		Length:  response.Length,
	}
	s.mapA.m[id] = proxyModel
	s.mapA.mu.Unlock()

	return proxyModel
}

func (s *StorageRepo) GetProxyData(id int) *model.ProxyModel {
	if value, ok := s.mapA.m[id]; ok {
		return &value
	}
	return nil
}
