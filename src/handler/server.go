package handler

import (
	"crawl/src/db"
	"crawl/src/util"
	"log"

	"github.com/robfig/cron/v3"
)

type Server struct {
	cfg     util.Config
	store   db.Store
	cronjob *cron.Cron
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	cronjob := cron.New()
	server := &Server{cfg: config, store: store, cronjob: cronjob}

	return server, nil
}

func (server *Server) Start() error {

	err := server.registerJob2()
	if err != nil {
		log.Println("can not register job 1")
	}
	server.cronjob.Start()
	var forever chan struct{}
	<-forever
	return nil
}
