package sqlstore

import (
	"database/sql"

	"github.com/ozon_test/internal/app/model"
	"github.com/ozon_test/internal/app/store"
)

type LinkRepository struct {
	store *Store
}

func (r *LinkRepository) IncreaseViews(l *model.Link) error {
	return r.store.db.QueryRow(
		"UPDATE links SET views = views + 1 WHERE token=($1) RETURNING views",
		l.Token,
	).Scan(&l.Views)
}
func (r *LinkRepository) Create(l *model.Link) error {
	err := l.BeforeCreate()
	if err != nil {
		return err
	}
	return r.store.db.QueryRow(
		"INSERT INTO links (link,token,created_at) VALUES ($1,$2,$3) RETURNING id",
		l.Link,
		l.Token,
		l.CreatedAt,
	).Scan(&l.ID)
}
func (r *LinkRepository) Find(token string) (*model.Link, error) {
	l := &model.Link{}
	if err := r.store.db.QueryRow(
		"SELECT id, link, token, created_at,views token FROM links WHERE token = $1",
		token,
	).Scan(
		&l.ID,
		&l.Link,
		&l.Token,
		&l.CreatedAt,
		&l.Views,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return l, nil
}
func (r *LinkRepository) FindByLink(link string) (*model.Link, error) {
	l := &model.Link{}
	if err := r.store.db.QueryRow(
		"SELECT id, link, token, created_at,views token FROM links WHERE link = $1",
		link,
	).Scan(
		&l.ID,
		&l.Link,
		&l.Token,
		&l.CreatedAt,
		&l.Views,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return l, nil
}
