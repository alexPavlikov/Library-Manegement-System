package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/alexPavlikov/Library-Manegement-System/internal/config"
	"github.com/alexPavlikov/Library-Manegement-System/internal/entity/author"
	"github.com/alexPavlikov/Library-Manegement-System/internal/entity/book"
	"github.com/alexPavlikov/Library-Manegement-System/internal/entity/user"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/client/postgresql"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func Run() {
	logger := logging.GetLogger()
	logger.Info("Create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	clientPostgreSQL, err := postgresql.NewClient(context.TODO(), cfg.Storage)
	if err != nil {
		logger.Fatalf("failed to get new client postgresql, due to err: %v", err)
	}

	logger.Info("Register book handlers")
	bookRep := book.NewRepository(clientPostgreSQL, logger)
	bookService := book.NewService(logger, bookRep)
	bookHandler := book.NewHandler(logger, bookService)
	bookHandler.Register(router)

	logger.Info("Register user handlers")
	userRep := user.NewRepository(clientPostgreSQL, logger)
	userService := user.NewService(logger, userRep)
	userHandler := user.NewHandler(logger, userService)
	userHandler.Register(router)

	logger.Info("Register author handlers")
	authorRep := author.NewRepository(clientPostgreSQL, logger)
	authorService := author.NewService(logger, authorRep)
	authorHandler := author.NewHandler(authorService, logger)
	authorHandler.Register(router)

	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("Start application")
	var listener net.Listener
	var listenErr error

	logger.Info("Listen TCP")
	listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	logger.Infof("Server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	if listenErr != nil {
		logger.Fatal(listenErr.Error())
	}
	server := &http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	err := server.Serve(listener)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
