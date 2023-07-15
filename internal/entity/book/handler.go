package book

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/alexPavlikov/Library-Manegement-System/internal/entity/genre"
	"github.com/alexPavlikov/Library-Manegement-System/internal/handlers"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/utils"
	"github.com/julienschmidt/httprouter"
)

var (
	Auth bool
	//Genres    []Genre
	Book_DTO DTO
	//	URL_NAME = []string{"Главная"}
	URL_MAP = map[string]string{"Главная": "/", "Книги": "/books/genre/all", "Жанры": "/books/genre/all", "Популярные": "/books/", "Авторы": "/authors/"}
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

	router.HandlerFunc(http.MethodGet, "/", h.IndexHandler)
	router.HandlerFunc(http.MethodGet, "/books/", h.PopularBooksHandler)
	router.HandlerFunc(http.MethodGet, "/books/genre/:uuid", h.GenreHandler)
	router.HandlerFunc(http.MethodGet, "/books/list/:uuid", h.BooksListHandler)
	//router.HandlerFunc(http.MethodGet, "/book/all", h.GetAllBookHandler)
	router.HandlerFunc(http.MethodGet, "/book/:uuid", h.BookHandler)

	router.HandlerFunc(http.MethodPost, "/book/comments/add", h.AddCommentHandler)
	router.HandlerFunc(http.MethodPost, "/book/download", h.BookDownloadHandler)

	router.HandlerFunc(http.MethodPost, "/books/find/", h.BookFindHandler)
	router.HandlerFunc(http.MethodGet, "/books/find/", h.GetBookFindHandler)

	router.HandlerFunc(http.MethodGet, "/test", h.test)
}

func (h *handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	Genres, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	books, err := h.service.GetNewBooks(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	URL_NAME := []string{"Главная"}
	page := map[string]interface{}{"Genres": Genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": "Новинки", "Auth": Book_DTO.Auth, "Books": books}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "index", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

}

func (h *handler) PopularBooksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	books, err := h.service.GetAllBooks(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	genres, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	URL_NAME := []string{"Главная", "Популярные"}

	page := map[string]interface{}{"Genres": genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": "Популярные", "Auth": Book_DTO.Auth, "Books": books}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "index", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}

func (h *handler) BooksListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	author := strings.TrimPrefix(r.URL.Path, "/books/list/")

	books, authorLastname, err := h.service.GetAllBooksByAuthor(context.TODO(), author)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	genres, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	URL_NAME := []string{"Главная", "Книги", authorLastname}

	page := map[string]interface{}{"Genres": genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": "Список книг", "Auth": Book_DTO.Auth, "Books": books}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "index", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}

func (h *handler) GenreHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	genresArr, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	link := strings.TrimPrefix(r.URL.Path, "/books/genre/")

	var books []Book
	var genres genre.Genre
	var URL_NAME []string

	if link == "all" {

		books, err = h.service.GetAllBooks(context.TODO())
		if err != nil {
			h.logger.Tracef("failed: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			//return err
		}
		URL_NAME = []string{"Главная", "Все жанры"}
		//genres.Name = "Все жанры"
		// else {
		// 	books, err = h.service.GetNewBooks(context.TODO())
		// 	if err != nil {
		//		h.logger.Tracef("failed: %v", err)
		// 		w.WriteHeader(http.StatusBadRequest)
		// 		//return err
		// 	}
		// 	URL_NAME = []string{"Главная", "Новинки"}
		// 	genre.Name = "Новинки"
		// }

	} else {
		genres, err = genre.GetGenreByLink(context.TODO(), link)
		if err != nil {
			h.logger.Tracef("failed: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			//return err
		}
		fmt.Println("GENRE!!!!!!!!!!!!!", genres)
		books, err = h.service.GetAllBooksByGenre(context.TODO(), genres.Id)
		if err != nil {
			h.logger.Tracef("failed: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			//return err
		}
		URL_NAME = []string{"Главная", "Жанры", genres.Name}
	}

	page := map[string]interface{}{"Genres": genresArr, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": genres.Name, "Auth": Book_DTO.Auth, "Books": books}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "index", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}

func (h *handler) GetAllBookHandler(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAllBooks(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	b, err := json.Marshal(books)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h *handler) BookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	uuid := strings.TrimPrefix(r.URL.Path, "/book/")
	fmt.Println(uuid)

	book, err := h.service.GetBook(context.TODO(), uuid)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	genres, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	URL_MAP[book.Name] = "/book/" + book.UUID

	URL_NAME := []string{"Главная", "Книги", book.Name}
	// var books []Book
	// books = append(books, book)

	comments, ok, err := h.service.GetAllCommentByBook(context.TODO(), book.UUID, Book_DTO.User_id)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	page := map[string]interface{}{"Genres": genres, "Books": book, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": book.Name, "Auth": Book_DTO.Auth, "User": Book_DTO.User_id, "Comments": comments, "ComOK": ok}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	err = tmpl.ExecuteTemplate(w, "book", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		//w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}

func (h *handler) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var comment Comment
	comment.User_id = r.FormValue("user")
	comment.Book_id = r.FormValue("book")
	comment.Text = r.FormValue("comment")
	comment.Time = time.Now().Format("2006-01-02 15:04")
	fmt.Println("COMMENT", comment)
	err := h.service.CreateCommentForBook(context.TODO(), &comment)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	http.Redirect(w, r, fmt.Sprintf("/book/%s", comment.Book_id), http.StatusSeeOther)
}

func (h *handler) BookDownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("uuid")
	book, err := h.service.GetBook(context.TODO(), text)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	path := "C:/Users/admin/go/src/Library-Manegement-System/assets/books/" + book.PDFLink
	userFile := "C:/Users/admin/Desktop/" + book.Name + ".pdf"
	err = utils.DownloadBook(path, userFile)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	http.Redirect(w, r, fmt.Sprintf("/book/%s", book.UUID), http.StatusSeeOther)
}

type R struct {
	Key  string `json:"Key"`
	Text string `json:"Text"`
}

var Books []Book
var rs R

func (h *handler) BookFindHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("response Body:", string(body))

	err = json.Unmarshal(body, &rs)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	Books, err = h.service.GetBookByName(context.TODO(), rs.Text)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

}

func (h *handler) GetBookFindHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	genres, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	URL_NAME := []string{"Главная", "Книги", fmt.Sprintf(`Результаты поиска по запросу "%s"`, rs.Text)}

	page := map[string]interface{}{"Genres": genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": fmt.Sprintf(`Результаты поиска по запросу "%s"`, rs.Text), "Auth": Book_DTO.Auth, "Books": Books}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "index", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}

var i = 1

func (h *handler) test(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	books, err := h.service.Test(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	r.ParseForm()
	id := r.FormValue("id")

	var class, clss string

	if id == "m" {

		if i > 1 {
			class = ""
			i -= 1

		} else {
			class = "untouchable"
		}
	} else if id == "p" {
		if len(books[i+1]) == 12 {
			i += 1
		} else {
			i += 1
			clss = "untouchable"
		}
	}

	Genres, err := genre.GetAllGenres(context.TODO())
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	URL_NAME := []string{"Главная", "Книги", "Test"}

	page := map[string]interface{}{"Genres": Genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Title": "Test", "Auth": Book_DTO.Auth, "Books": books[i], "Index": i, "Class1": class, "Class2": clss}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "test", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}
