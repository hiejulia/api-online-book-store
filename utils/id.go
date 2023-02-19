package utils

import uuid "github.com/satori/go.uuid"

func ID() string {
	return uuid.NewV4().String()
}
