package internalstore_test

import (
	"testing"

	"github.com/ozon_test/internal/app/model"
	"github.com/ozon_test/internal/app/store"
	"github.com/ozon_test/internal/app/store/internalstore"
	"github.com/stretchr/testify/assert"
)

func TestLinkRepository_IncreaseViews(t *testing.T) {
	l := model.TestLink(t)

	s := internalstore.New()
	s.Link().Create(l)
	oldViews := l.Views

	err := s.Link().IncreaseViews(l)
	assert.NoError(t, err)
	assert.Equal(t, oldViews+1, l.Views)
}

func TestLinkRepository_Create(t *testing.T) {
	l := model.TestLink(t)

	s := internalstore.New()
	err := s.Link().Create(l)
	assert.NoError(t, err)

	err = s.Link().Create(l)
	assert.Error(t, err)

	l, err = s.Link().Find(l.Token)
	assert.NoError(t, err)
	assert.NotEqual(t, l.ID, 0)
}
func TestLinkRepository_Find(t *testing.T) {
	l := model.TestLink(t)

	s := internalstore.New()
	s.Link().Create(l)

	newL, err := s.Link().Find(l.Token)
	assert.NoError(t, err)
	assert.Equal(t, newL.ID, l.ID)

	newL, err = s.Link().Find("abc")
	assert.Error(t, err)
	assert.Equal(t, err, store.ErrRecordNotFound)
}

func TestLinkRepository_FindByLink(t *testing.T) {
	l := model.TestLink(t)

	s := internalstore.New()
	s.Link().Create(l)

	newL, err := s.Link().FindByLink(l.Link)
	assert.NoError(t, err)
	assert.Equal(t, newL.ID, l.ID)

	newL, err = s.Link().FindByLink("wrong.com")
	assert.Error(t, err)
	assert.Equal(t, err, store.ErrRecordNotFound)
}
