package main

import "github.com/gorilla/mux"

func (srv *server) Routes() {
	handler := mux.NewRouter()
	handler.HandleFunc("/", srv.logRequest(srv.HandleTemplate("templates/index.html"))).Methods("GET")
	handler.HandleFunc("/scripts/index.js", srv.logRequest(srv.ServeFile("templates/scripts/index.js"))).Methods("GET")
	handler.HandleFunc("/stylesheets/index.css", srv.logRequest(srv.ServeFile("templates/stylesheets/index.css"))).Methods("GET")
	handler.HandleFunc("/stylesheets/minireset.css", srv.logRequest(srv.ServeFile("templates/stylesheets/minireset.css"))).Methods("GET")
	srv.Server.Handler = handler
}
