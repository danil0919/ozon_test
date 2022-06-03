package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

var (
	testDatabaseURL string = "host=localhost dbname=test_ozon_test sslmode=disable"
)

// TestStore
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}
		db.Close()
	}

}

func TestStore(t *testing.T) (*Store, func(...string)) {
	db, teardown := TestDB(t, testDatabaseURL)

	s := New(db)
	return s, teardown
}
