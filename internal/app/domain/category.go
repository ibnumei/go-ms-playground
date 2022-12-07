package domain

import "time"

type Category struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"categoryName"`
	IsActive     bool      `json:"isActive"`
	CreatedBy    string    `json:"createdBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedBy    string    `json:"updatedBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
