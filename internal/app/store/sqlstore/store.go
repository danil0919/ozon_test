package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ozon_test/internal/app/store"
)

type Store struct {
	db             *sql.DB
	linkRepository *LinkRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//User....
func (s *Store) Link() store.LinkRepository {
	if s.linkRepository != nil {
		return s.linkRepository
	}

	s.linkRepository = &LinkRepository{
		store: s,
	}

	return s.linkRepository
}

// store.User().Create()
