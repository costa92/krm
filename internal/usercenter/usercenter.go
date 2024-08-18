package usercenter

import (
	"github.com/costa92/krm/pkg/log"
	"github.com/go-kratos/kratos/v2"
)

// Server represents the server.
type Server struct {
	app *kratos.App
}

// Run is a method of the Server struct that starts the server.
func (s *Server) Run(stopCh <-chan struct{}) error {
	go func() {
		if err := s.app.Run(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-stopCh

	log.Infof("Gracefully shutting down server ...")

	if err := s.app.Stop(); err != nil {
		log.Errorw(err, "Failed to gracefully shutdown kratos application")
		return err
	}

	return nil
}
