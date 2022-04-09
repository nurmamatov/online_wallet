package handlers

import (
	"online_wallet/config"
	"online_wallet/storage"
)

type Handler struct {
	cfg  config.Config
	repo storage.Storage
}

func NewHandler(cfg config.Config, storage storage.Storage) *Handler {
	return &Handler{
		cfg:  cfg,
		repo: storage,
	}
}