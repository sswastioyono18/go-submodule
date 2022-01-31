package tools_1

import "github.com/google/uuid"

func NewGoogleUUID() string {
	return uuid.NewString()
}
