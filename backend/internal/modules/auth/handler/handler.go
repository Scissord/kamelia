package handler

import (
	"backend/internal/modules/auth/service"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{service: s}
}
