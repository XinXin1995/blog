package utils

import (
	uuid "github.com/satori/go.uuid"
)

func CreateUuid() string {
	return uuid.NewV4().String()
}
