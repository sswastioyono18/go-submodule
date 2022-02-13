package tools_2

import (
	"fmt"
	"github.com/gofrs/uuid"
)

func NewGofrsUUID() string {
	strResult, _ := uuid.NewV4()
	fmt.Println("TEST123")
	return strResult.String()
}
