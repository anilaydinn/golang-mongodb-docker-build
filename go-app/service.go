package main

import (
	"strings"

	"github.com/google/uuid"
)

type Data struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Service struct {
	repository *Repository
}

const SecretKey = "14465375-b4a8-47fa-9692-c986d4a825ee"

func NewService(repository *Repository) Service {
	return Service{
		repository: repository,
	}
}

func GenerateUUID(length int) string {
	uuid := uuid.New().String()

	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}
