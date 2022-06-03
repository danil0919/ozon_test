package model

import (
	"time"

	"github.com/ozon_test/internal/app/shortener"
)

type Link struct {
	ID        int
	Link      string
	Token     string
	Views     int
	CreatedAt time.Time
}

func (l *Link) BeforeCreate() error {
	token, err := shortener.GenerateShortLink(l.Link)
	if err != nil {
		return err
	}
	l.Token = token

	l.CreatedAt = time.Now()
	return nil
}
func (l *Link) BeforeCreateExpiresAt(expires time.Time) error {
	err := l.BeforeCreate()
	if err != nil {
		return err
	}
	l.CreatedAt = expires
	return nil
}
