package repository

import "github.com/roxyash/kmf_testtask/proxy_service/internal/repository/model"

type Storage interface {
	SetProxyData(response model.ProxyModel) model.ProxyModel
	GetProxyData(id int) *model.ProxyModel
}

type Repository struct {
	Storage
}

func NewRepository() *Repository {
	return &Repository{
		Storage: NewSomeStorageRepo(),
	}
}
