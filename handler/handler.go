package handler

import "github.com/google/uuid"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func CreateNewId() string {
	return uuid.New().String()
}
