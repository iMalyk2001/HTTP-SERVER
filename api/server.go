package api

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Items struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}


type Server struct {
	*mux.Router
	shoppingList []Items
}

func NewServer() *Server {
	s := &Server{
		Router: 	   mux.NewRouter(),
		shoppingList:  []Items{},
	}
	return s

}