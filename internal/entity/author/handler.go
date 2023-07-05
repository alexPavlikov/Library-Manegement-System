package author

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/alexPavlikov/Library-Manegement-System/internal/entity/book"
	"github.com/alexPavlikov/Library-Manegement-System/internal/entity/genre"
	"github.com/alexPavlikov/Library-Manegement-System/internal/handlers"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

var (
	URL_MAP = map[string]string{"Главная": "/", "Книги": "/books/genre/all", "Жанры": "/books/genre/all", "Популярные": "/books/", "Авторы": "/authors/"}
)

type handler struct {
	service *Service
	logger  *logging.Logger
}

// Register implements handlers.Handlers.
func (h *handler) Register(router *httprouter.Router) {

	router.HandlerFunc(http.MethodGet, "/authors/", h.GetAuthorsHandler)
	router.HandlerFunc(http.MethodPost, "/authors/find/", h.AuthorFindHandler)
	router.HandlerFunc(http.MethodGet, "/authors/find/", h.GetAuthorsFindHandler)
	router.HandlerFunc(http.MethodGet, "/author/:uuid", h.GetAuthorHandler)
}

func NewHandler(service *Service, logger *logging.Logger) handlers.Handlers {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	authors, err := h.service.GetAllAuthors(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	genres, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	URL_NAME := []string{"Главная", "Авторы"}
	page := map[string]interface{}{"Genres": genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": "Новинки", "Auth": book.Book_DTO.Auth, "Authors": authors}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = tmpl.ExecuteTemplate(w, "authors", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

type R struct {
	Key  string `json:"Key"`
	Text string `json:"Text"`
}

var Authors []Author
var rs R

func (h *handler) AuthorFindHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &rs)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	Authors, err = h.service.GetAuthorByName(context.TODO(), rs.Text)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *handler) GetAuthorsFindHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	genres, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	URL_NAME := []string{"Главная", "Авторы", fmt.Sprintf(`Результаты поиска по запросу "%s"`, rs.Text)}

	page := map[string]interface{}{"Genres": genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": "Новинки", "Auth": book.Book_DTO.Auth, "Authors": Authors}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = tmpl.ExecuteTemplate(w, "authors", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *handler) GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	uuid := strings.TrimPrefix(r.URL.Path, "/author/")
	author, err := h.service.GetAuthor(context.TODO(), uuid)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	text, err := h.service.FindBiography(context.TODO(), author.UUID)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	genres, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	fullName := author.Firstname + " " + author.Lastname

	URL_NAME := []string{"Главная", "Авторы", fullName}

	page := map[string]interface{}{"Genres": genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": fullName, "Auth": book.Book_DTO.Auth, "Author": author, "Text": text}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = tmpl.ExecuteTemplate(w, "author_biography", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
