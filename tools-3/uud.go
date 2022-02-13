package tools_3

import (
	"fmt"
	"github.com/gofrs/uuid"
)

func NewGofrsUUID() string {
	strResult, _ := uuid.NewV4()
	fmt.Println("test tools-3 berubah")
	return strResult.String()
}
