package models

import "github.com/google/uuid"

type Link struct {
	Id  uuid.UUID `json:"id"`
	Url string    `json:"url"`
}
