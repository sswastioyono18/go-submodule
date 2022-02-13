package tools_2

import "github.com/gofrs/uuid"

func NewGofrsUUID() string {
	strResult, _ := uuid.NewV4()
	return strResult.String()
}
