package entity

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	ID        uuid.UUID
	Nim       string
	Name      string
	Class     Class
	Team      Team
	CreatedAt time.Time
	UpdatedAt time.Time
}
