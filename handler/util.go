package handler

import "github.com/google/uuid"

func randomID() string {
	return uuid.New().String()
}
