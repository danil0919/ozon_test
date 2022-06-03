package model_test

import (
	"testing"

	"github.com/ozon_test/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestLink_BeforeCreate(t *testing.T) {
	l := model.TestLink(t)

	l.BeforeCreate()
	assert.NotEqual(t, l.Token, "")
	assert.Equal(t, len(l.Token), 10)
	assert.NotEmpty(t, l.CreatedAt)
}
