package user

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/alexPavlikov/Library-Manegement-System/internal/entity/book"
	"github.com/alexPavlikov/Library-Manegement-System/internal/handlers"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

var (
	URL_MAP = map[string]string{"Главная": "/books", "Новинки": "/books/new", "Книги": "/books/all", "Вход/Регистрация": "/user/"}
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
	//router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	router.HandlerFunc(http.MethodGet, "/user/", h.UserProfileHandler)
	router.HandlerFunc(http.MethodPost, "/user/reg", h.UserRegHandler)
	router.HandlerFunc(http.MethodPost, "/user/auth", h.UserAuthHandler)
}

func (h *handler) UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./internal/html/*.html")
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

	URL_NAME := []string{"Главная", "Вход | Регистрация"}

	page := map[string]interface{}{"Genres": book.GenresDTO.Genres, "URLs": URL_MAP, "URL_NAME": URL_NAME, "Auth": false, "Auth_title": "Войти в аккаунт", "Reg_title": "Регистрация"}

	err = tmpl.ExecuteTemplate(w, "header", nil)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "authreg", page)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}

}

func (h *handler) UserRegHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := User{
		Id:           "",
		Firstname:    r.FormValue("fname"),
		Lastname:     r.FormValue("lname"),
		Age:          0,
		DateOfBirth:  r.FormValue("date"),
		Gender:       r.FormValue("gender"),
		Login:        r.FormValue("login"),
		PasswordHash: r.FormValue("password"),
	}
	fmt.Println(user)
	err := h.service.CreateUser(context.TODO(), &user)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (h *handler) UserAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	login := r.FormValue("login")
	password := r.FormValue("password")

	user, err := h.service.GetAuth(context.TODO(), login, password)
	if err != nil {
		h.logger.Tracef("failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(user)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
