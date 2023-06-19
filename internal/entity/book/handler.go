package book

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/alexPavlikov/Library-Manegement-System/internal/handlers"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

var (
	Auth bool
	//	URL_NAME = []string{"Главная"}
	URL_MAP = map[string]string{"Главная": "/books", "Новинки": "/books/new", "Книги": "/books/all"}
)

type handler struct {
	logger  *logging.Logger
	service *Service
}

func NewHandler(logger *logging.Logger, service *Service) handlers.Handlers {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	router.HandlerFunc(http.MethodGet, "/books/", h.IndexHandler)
	router.HandlerFunc(http.MethodGet, "/books/genre/:uuid", h.GenreHandler)
	//router.HandlerFunc(http.MethodGet, "/book/all", h.GetAllBookHandler)
	router.HandlerFunc(http.MethodGet, "/book/:uuid", h.BookHandler)

	// router.HandlerFunc(http.MethodPost, "/book/comments/add", h.AddCommentHandler)
	// router.HandlerFunc(http.MethodPost, "/book/download", h.BookDownloadHandler)

	router.HandlerFunc(http.MethodGet, "/test", h.TestHandler)
}

func (h *handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	genres, err := h.service.GetAllGenres(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	books, err := h.service.GetAllBooks(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	URL_NAME := []string{"Главная"}

	page := map[string]interface{}{"Genres": genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": "Популярные", "Auth": false, "Books": books}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "index", page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}

func (h *handler) GenreHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	genres, err := h.service.GetAllGenres(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	link := strings.TrimPrefix(r.URL.Path, "/books/genre/")

	var books []Book
	var genre Genre
	var URL_NAME []string

	if link == "all" || link == "new" {

		if link == "all" {
			books, err = h.service.GetAllBooks(context.TODO())
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				//return err
			}
			URL_NAME = []string{"Главная", "Все жанры"}
			genre.Name = "Все жанры"
		} else {
			books, err = h.service.GetAllBooks(context.TODO())
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				//return err
			}
			URL_NAME = []string{"Главная", "Новинки"}
			genre.Name = "Новинки"
		}

	} else {
		genre, err = h.service.GetGenreByLink(context.TODO(), link)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			//return err
		}

		books, err = h.service.GetAllBooksByGenre(context.TODO(), genre.Id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			//return err
		}
		URL_NAME = []string{"Главная", "Жанры", genre.Name}
	}

	page := map[string]interface{}{"Genres": genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": genre.Name, "Auth": false, "Books": books}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "index", page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}

func (h *handler) GetAllBookHandler(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAllBooks(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	b, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h *handler) BookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	uuid := strings.TrimPrefix(r.URL.Path, "/book/")
	fmt.Println(uuid)

	book, err := h.service.GetBook(context.TODO(), uuid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	fmt.Println(book)
	genres, err := h.service.GetAllGenres(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	URL_MAP[book.Name] = "/book/" + book.UUID

	URL_NAME := [3]string{"Главная", "Книги", book.Name}
	fmt.Println(URL_NAME)
	var books []Book
	books = append(books, book)

	page := map[string]interface{}{"Genres": genres, "Books": books, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": book.Name, "Auth": true}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	err = tmpl.ExecuteTemplate(w, "book", page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}

func (h *handler) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("comment")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(text))
}

func (*handler) BookDownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("uuid")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(text))
}

func (h *handler) TestHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	err = tmpl.ExecuteTemplate(w, "test", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}
