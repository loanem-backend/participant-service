package entity

import "time"

type Class struct {
	ID        int
	Name      string
	Course    Course
	CreatedAt time.Time
	UpdatedAt time.Time
}
