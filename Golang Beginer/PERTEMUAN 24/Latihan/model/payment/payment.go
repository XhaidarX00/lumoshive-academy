package payment

import (
	"database/sql"
	"time"
)

type Payment struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Photo     sql.NullString `json:"photo"`
	IsActive  bool           `json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at,omitempty"`
}
