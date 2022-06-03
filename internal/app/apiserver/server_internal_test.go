package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ozon_test/internal/app/model"
	"github.com/ozon_test/internal/app/store/internalstore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleLinksGet(t *testing.T) {
	store := internalstore.TestStore(make(map[string]*model.Link))
	s := NewServer(store)

	l := model.TestLink(t)
	s.store.Link().Create(l)
	testCases := []struct {
		name         string
		token        string
		expectedCode int
	}{
		{
			name:         "valid",
			token:        l.Token,
			expectedCode: http.StatusOK,
		},
		{
			name:         "empty token",
			token:        "",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "not found",
			token:        "aaa",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/"+tc.token, nil)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
func TestServer_HandleLinksCreate(t *testing.T) {
	store := internalstore.TestStore(make(map[string]*model.Link))
	s := NewServer(store)

	link := model.TestLink(t)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"long_url": link.Link,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "already exists",
			payload: map[string]string{
				"long_url": link.Link,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_HandleLinksGetInfo(t *testing.T) {
	store := internalstore.TestStore(make(map[string]*model.Link))
	s := NewServer(store)

	l := model.TestLink(t)
	s.store.Link().Create(l)
	testCases := []struct {
		name         string
		token        string
		expectedCode int
	}{
		{
			name:         "valid",
			token:        l.Token,
			expectedCode: http.StatusOK,
		},
		{
			name:         "empty token",
			token:        "",
			expectedCode: http.StatusNotFound, //gorilla returns 404 if url does not match pattern
		},
		{
			name:         "not found",
			token:        "aaa",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/info/"+tc.token, nil)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
