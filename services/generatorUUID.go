package services

import "github.com/google/uuid"

func NewUUID() string {

	uuid := uuid.New()
	return uuid.String()

}

