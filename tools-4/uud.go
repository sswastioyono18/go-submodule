package tools_4

import (
	"fmt"
	"github.com/gofrs/uuid"
)

func NewGofrsUUID() string {
	strResult, _ := uuid.NewV4()
	fmt.Println("TEST123 berubah lagi")
	return strResult.String()
}
