package models

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int64     `json:"age"`
}
