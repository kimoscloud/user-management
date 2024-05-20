package response

import (
	"time"
)

type UserLightDTO struct {
	ID        string    `json:"id"`
	LastName  string    `json:"lastName"`
	FirstName string    `json:"firstName"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	LastLogin time.Time `json:"lastLogin"`
	DeletedAt time.Time `json:"deletedAt"`
}
