package internalstore

import (
	"errors"
	"sync"

	"github.com/ozon_test/internal/app/model"
	"github.com/ozon_test/internal/app/store"
)

type LinkRepository struct {
	store     *Store
	mu        *sync.RWMutex
	links     map[string]*model.Link
	linksLong map[string]string
	lastID    int
}

func (r *LinkRepository) IncreaseViews(l *model.Link) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	l, ok := r.links[l.Token]
	if !ok {
		return store.ErrRecordNotFound
	}

	l.Views++
	return nil
}
func (r *LinkRepository) Create(l *model.Link) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.linksLong[l.Link]; ok {
		return errors.New("already exists")
	}
	err := l.BeforeCreate()
	if err != nil {
		return err
	}
	l.ID = r.lastID + 1
	r.lastID++
	r.links[l.Token] = l
	r.linksLong[l.Link] = l.Token
	return nil
}
func (r *LinkRepository) Find(token string) (*model.Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	l, ok := r.links[token]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return l, nil
}

func (r *LinkRepository) FindByLink(link string) (*model.Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if token, ok := r.linksLong[link]; ok {
		return r.links[token], nil
	}
	return nil, store.ErrRecordNotFound
}
