package book

import (
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
	router.HandlerFunc(http.MethodGet, "/book", h.indexHandler)
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
