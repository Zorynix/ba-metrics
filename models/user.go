package models

type User struct {
	Name  string `json:"name"`
	Count int    `json:"count,omitempty"`
}
