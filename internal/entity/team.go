package entity

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID        uuid.UUID
	Number    int
	Class     Class
	CreatedAt time.Time
	UpdatedAt time.Time
}
