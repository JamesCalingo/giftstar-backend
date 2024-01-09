package models

import "time"

type Item struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Link      string    `json:"link"`
	Claimed   bool      `json:"claimed"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
