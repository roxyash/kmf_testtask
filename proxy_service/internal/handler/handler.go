package handler

import (
	"github.com/roxyash/kmf_testtask/pkg/zaplogger"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/service"
)

type Handler struct {
	logger  zaplogger.Logger
	service *service.Service
}

func NewHandler(logger zaplogger.Logger, service *service.Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}
