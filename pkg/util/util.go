package util

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id" swaggertype:"string"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type APIRequest struct {
	Headers map[string]string
	Body    []byte
	Url     string
	Method  string
}

type APIResponse struct {
	StatusCode int
	Status     string
	Body       []byte
}
