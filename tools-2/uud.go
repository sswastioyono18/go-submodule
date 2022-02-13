package tools_2

import (
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

func NewGofrsUUID() string {
	strResult, _ := uuid.NewV4()
	log.Logger.Info().Msg("Test123")
	return strResult.String()
}
