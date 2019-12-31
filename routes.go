package main

import "github.com/gorilla/mux"

func (srv *server) Routes() {
	handler := mux.NewRouter()
	handler.HandleFunc("/", srv.LogRequest(srv.HandleTemplate("templates/index.html"))).Methods("GET")
	handler.HandleFunc("/motorcycles", srv.LogRequest(srv.HandleTemplate("templates/motorcycles.html"))).Methods("GET")
	handler.HandleFunc("/scripts/index.js", srv.LogRequest(srv.ServeFile("templates/scripts/index.js"))).Methods("GET")
	handler.HandleFunc("/stylesheets/index.css", srv.LogRequest(srv.ServeFile("templates/stylesheets/index.css"))).Methods("GET")
	handler.HandleFunc("/stylesheets/minireset.css", srv.LogRequest(srv.ServeFile("templates/stylesheets/minireset.css"))).Methods("GET")
	handler.HandleFunc("/scripts/motorcycles.js", srv.LogRequest(srv.ServeFile("templates/scripts/motorcycles.js"))).Methods("GET")
	handler.HandleFunc("/stylesheets/motorcycles.css", srv.LogRequest(srv.ServeFile("templates/stylesheets/motorcycles.css"))).Methods("GET")
	srv.Server.Handler = handler
}
