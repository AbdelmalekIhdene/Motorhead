package main

import "github.com/gorilla/mux"

func (srv *server) Routes() {
	handler := mux.NewRouter()
	handler.HandleFunc("/", srv.LogRequest(srv.HandleTemplate("templates/index.html"))).Methods("GET")
	handler.HandleFunc("/selection", srv.LogRequest(srv.HandleTemplate("templates/selection.html"))).Methods("GET")
	handler.HandleFunc("/scripts/index.js", srv.LogRequest(srv.ServeFile("templates/scripts/index.js"))).Methods("GET")
	handler.HandleFunc("/stylesheets/index.css", srv.LogRequest(srv.ServeFile("templates/stylesheets/index.css"))).Methods("GET")
	handler.HandleFunc("/stylesheets/minireset.css", srv.LogRequest(srv.ServeFile("templates/stylesheets/minireset.css"))).Methods("GET")
	handler.HandleFunc("/scripts/selection.js", srv.LogRequest(srv.ServeFile("templates/scripts/selection.js"))).Methods("GET")
	handler.HandleFunc("/stylesheets/selection.css", srv.LogRequest(srv.ServeFile("templates/stylesheets/selection.css"))).Methods("GET")
	srv.Server.Handler = handler
}
