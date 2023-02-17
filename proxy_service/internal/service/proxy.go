package service

import (
	"github.com/roxyash/kmf_testtask/proxy_service/internal/repository"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/repository/model"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/response"
)

type ProxyService struct {
	repository *repository.Repository
}

func NewProxyService(repository *repository.Repository) *ProxyService {
	return &ProxyService{
		repository: repository,
	}
}

func (s *ProxyService) SetProxyResponseData(rsp response.ProxyResponse) response.ProxyResponse {
	proxyModel := s.repository.SetProxyData(model.ProxyModel{
		ID:      rsp.ID,
		Status:  rsp.Status,
		Headers: rsp.Headers,
		Length:  rsp.Length,
	})

	return response.ProxyResponse{
		ID:      proxyModel.ID,
		Status:  proxyModel.Status,
		Headers: proxyModel.Headers,
		Length:  proxyModel.Length,
	}
}
