package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ozon_test/internal/app/model"
	"github.com/ozon_test/internal/app/store"
)

func (s *server) handleLinkGet() http.HandlerFunc {
	type response struct {
		Link     string `json:"long_url,omitempty"`
		ShortUrl string `json:"short_url,omitempty"`
		Views    int    `json:"views,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		t := mux.Vars(r)["token"]
		link, err := s.store.Link().Find(t)
		if err != nil {
			s.logger.Debug(err)
			s.error(w, r, http.StatusNotFound, store.ErrRecordNotFound)
			return
		}
		res := &response{
			Link: link.Link,
		}
		s.respond(w, r, http.StatusOK, res)
		s.store.Link().IncreaseViews(link)
	}
}
func (s *server) handleLinkCreate() http.HandlerFunc {
	type request struct {
		Link string `json:"long_url"`
	}
	type response struct {
		Link     string `json:"long_url,omitempty"`
		ShortUrl string `json:"short_url,omitempty"`
		Message  string `json:"message,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Debug(err)
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		link := &model.Link{
			Link: req.Link,
		}

		err := s.store.Link().Create(link)
		res := &response{}
		if err != nil {
			linkExist, err2 := s.store.Link().FindByLink(link.Link)
			if err2 != nil {
				s.logger.Debug("Error creating short link", err, err2)
				s.error(w, r, http.StatusInternalServerError, errors.New("Internal error, please try again!"))
				return
			}
			res.Link = linkExist.Link
			res.ShortUrl = Host + "/" + linkExist.Token
			res.Message = "Short url for this url already exists"
			s.respond(w, r, http.StatusOK, res)
			return
		}

		res.Link = link.Link
		res.ShortUrl = Host + "/" + link.Token
		res.Message = "Short link was successfully created"
		s.respond(w, r, http.StatusCreated, res)
	}
}

func (s *server) handleLinkGetInfo() http.HandlerFunc {
	type response struct {
		Link      string    `json:"long_url,omitempty"`
		ShortUrl  string    `json:"short_url,omitempty"`
		Views     int       `json:"views,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		t := mux.Vars(r)["token"]
		link, err := s.store.Link().Find(t)
		if err != nil {
			s.logger.Debug(err)
			s.error(w, r, http.StatusNotFound, store.ErrRecordNotFound)
			return
		}

		res := &response{
			Link:      link.Link,
			ShortUrl:  Host + "/" + link.Token,
			Views:     link.Views,
			CreatedAt: link.CreatedAt,
		}
		s.respond(w, r, http.StatusOK, res)
	}
}
