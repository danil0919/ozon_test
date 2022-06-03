package internalstore

import (
	"sync"

	"github.com/ozon_test/internal/app/model"
	"github.com/ozon_test/internal/app/store"
)

type Store struct {
	loadLinkRepositoryOnce sync.Once
	linkRepository         *LinkRepository
}

func New() *Store {
	return &Store{}
}

//User....
func (s *Store) Link() store.LinkRepository {
	if s.linkRepository != nil {
		return s.linkRepository
	}
	s.loadLinkRepositoryOnce.Do(func() {
		s.linkRepository = &LinkRepository{
			store:     s,
			links:     make(map[string]*model.Link),
			linksLong: make(map[string]string),
			mu:        &sync.RWMutex{},
		}
	})

	return s.linkRepository
}
