package store

import "github.com/ozon_test/internal/app/model"

type LinkRepository interface {
	Create(*model.Link) error
	Find(string) (*model.Link, error)
	IncreaseViews(*model.Link) error
	FindByLink(string) (*model.Link, error)
}
