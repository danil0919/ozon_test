package sqlstore_test

import (
	"testing"

	"github.com/ozon_test/internal/app/model"
	"github.com/ozon_test/internal/app/store"
	"github.com/ozon_test/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestLinkRepository_IncreaseViews(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("links")
	s := sqlstore.New(db)
	l := model.TestLink(t)
	s.Link().Create(l)

	s.Link().IncreaseViews(l)
	assert.Equal(t, 1, l.Views)
}
func TestLinkRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("links")

	s := sqlstore.New(db)
	l := model.TestLink(t)
	err := s.Link().Create(l)

	assert.NoError(t, err)
	assert.NotNil(t, l)
	err = s.Link().Create(l)
	assert.Error(t, err)
}
func TestLinkRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("links")

	s := sqlstore.New(db)
	l := model.TestLink(t)
	s.Link().Create(l)

	l, err := s.Link().Find(l.Token)
	assert.NoError(t, err)
	assert.NotNil(t, l)

	l, err = s.Link().Find("abc")

	assert.Nil(t, l)
	assert.Error(t, err)
	assert.Equal(t, store.ErrRecordNotFound, err)
}

func TestLinkRepository_FindByLink(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("links")

	s := sqlstore.New(db)
	l := model.TestLink(t)
	l.BeforeCreate()
	s.Link().Create(l)

	l, err := s.Link().FindByLink(l.Link)
	assert.NoError(t, err)
	assert.NotNil(t, l)

	l, err = s.Link().FindByLink("wrong.com")

	assert.Nil(t, l)
	assert.Error(t, err)
	assert.Equal(t, store.ErrRecordNotFound, err)
}
