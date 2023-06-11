package book

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/alexPavlikov/Library-Manegement-System/internal/handlers"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
	"github.com/julienschmidt/httprouter"
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
	router.HandlerFunc(http.MethodGet, "/book/", h.indexHandler)
	router.HandlerFunc(http.MethodGet, "/book/all", h.GetAllBookHandler)
	router.HandlerFunc(http.MethodPost, "/book/", h.BookHandler)
}

func (h *handler) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./internal/html/index.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	err = tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
}

func (h *handler) GetAllBookHandler(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAll(context.TODO())
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
	r.ParseForm()
	id := r.FormValue("id")
	book, err := h.service.GetBook(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	b, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
