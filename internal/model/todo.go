package model

import "time"

type Todo struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}