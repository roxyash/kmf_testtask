package service

import (
	"github.com/roxyash/kmf_testtask/proxy_service/internal/repository"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/response"
)

type Proxy interface {
	SetProxyResponseData(response response.ProxyResponse) response.ProxyResponse
}

type Service struct {
	Proxy
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Proxy: NewProxyService(repository),
	}
}
