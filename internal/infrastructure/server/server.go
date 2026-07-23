package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	AppAddress *string
	AppPort    *string
	Engine     *gin.Engine
}

func NewServer(appAddress *string, appPort *string, engine *gin.Engine) *Server {
	return &Server{
		AppAddress: appAddress,
		AppPort:    appPort,
		Engine:     engine,
	}
}

func (s *Server) Run() *http.Server {
	addressServer := *s.AppAddress + *s.AppPort

	server := &http.Server{
		Addr:    addressServer,
		Handler: s.Engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			messageError := fmt.Errorf("server.go(Run) listen: %w", err)
			log.Fatalln(messageError)
		}
	}()

	log.Printf("Address: %s\n", addressServer)

	return server
}

func (s *Server) Close(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		messageError := fmt.Errorf("server.go(Close) Server forced to shutdown: %w", err)
		log.Fatal(messageError)
	}

	log.Println("Server exiting")
}
