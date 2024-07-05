package check

import (
	"time"

	"github.com/google/uuid"
)

type CheckResponse struct {
	ID        uuid.UUID `db:"id"`
	Code      string    `db:"code"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
