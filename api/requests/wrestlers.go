package requests

import (
	"github.com/google/uuid"
)

type CreateWrestlerRequest struct {
	Name string `json:"name" binding:"required"`
}

type GetWrestlerRequest struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
