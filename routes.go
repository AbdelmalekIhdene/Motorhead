package main

import "github.com/gorilla/mux"

func (srv *server) Routes() {
	handler := mux.NewRouter()
	handler.HandleFunc("/", srv.LogRequest(srv.HandleTemplate("static/index.html"))).Methods("GET")
	handler.HandleFunc("/motorcycles", srv.LogRequest(srv.HandleTemplate("static/motorcycles.html"))).Methods("GET")
	handler.HandleFunc("/scripts/{filename}", srv.LogRequest(srv.ServeDirectory("static/scripts"))).Methods("GET")
	handler.HandleFunc("/stylesheets/{filename}", srv.LogRequest(srv.ServeDirectory("static/stylesheets"))).Methods("GET")
	handler.HandleFunc("/images/hero/{filename}", srv.LogRequest(srv.ServeDirectory("static/images/hero"))).Methods("GET")
	handler.HandleFunc("/images/sideview/{filename}", srv.LogRequest(srv.ServeDirectory("static/images/sideview"))).Methods("GET")
	srv.Server.Handler = handler
}
