package internalstore

import (
	"sync"

	"github.com/ozon_test/internal/app/model"
)

func TestStore(testData map[string]*model.Link) *Store {
	s := &Store{}
	return &Store{
		linkRepository: &LinkRepository{
			store: s,
			links: testData,
			mu:    &sync.RWMutex{},
		},
	}
}
