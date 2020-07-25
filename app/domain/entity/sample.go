package entity

import (
	"time"
)

type Sample struct {
	ID        int `gorm:"primary_key"`
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Sample) TableName() string {
	return "samples"
}

func NewSample(id int, title string) *Sample {
	now := time.Now()
	return &Sample{
		Title: title, CreatedAt: now, UpdatedAt: now,
	}
}

func (s *Sample) GetID() int {
	return s.ID
}

func (s *Sample) GetTitle() string {
	return s.Title
}

func (s *Sample) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *Sample) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}