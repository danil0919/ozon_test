package apiserver

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/ozon_test/internal/app/store"
	"github.com/ozon_test/internal/app/store/internalstore"
	"github.com/ozon_test/internal/app/store/sqlstore"
)

var (
	Host string
)

func Start(config *Config) error {
	Host = config.Host
	var store store.Store
	switch config.StoreType {
	case "sql":
		db, err := newDB(config.DatabaseURL)
		if err != nil {
			return err
		}
		defer db.Close()
		store = sqlstore.New(db)
	case "internal":
		store = internalstore.New()
	default:
		return errors.New("undefined store type")
	}
	srv := NewServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(database_url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil

}
