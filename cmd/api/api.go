package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	user "github.com/yaviral17/go-ecom/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening to port 8080")

	return http.ListenAndServe(s.addr, router)
}
