package main

import "github.com/gorilla/mux"

func (srv *server) Routes() {
	handler := mux.NewRouter()
	handler.HandleFunc("/", srv.LogRequest(srv.HandleStaticTemplate("static/index.html"))).Methods("GET")
	handler.HandleFunc("/selection", srv.LogRequest(srv.HandleSelectionTemplate("static/selection.html"))).Methods("GET")
	handler.HandleFunc("/motorcycle", srv.LogRequest(srv.HandleMotorcycleTemplate("static/motorcycle.html"))).Methods("GET")
	handler.HandleFunc("/scripts/{filename}", srv.LogRequest(srv.ServeDirectory("static/scripts"))).Methods("GET")
	handler.HandleFunc("/components/{filename}", srv.LogRequest(srv.ServeDirectory("static/components"))).Methods("GET")
	handler.HandleFunc("/stylesheets/{filename}", srv.LogRequest(srv.ServeDirectory("static/stylesheets"))).Methods("GET")
	handler.HandleFunc("/images/hero/{filename}", srv.LogRequest(srv.ServeDirectory("static/images/hero"))).Methods("GET")
	handler.HandleFunc("/images/sideview/{filename}", srv.LogRequest(srv.ServeDirectory("static/images/sideview"))).Methods("GET")
	srv.Server.Handler = handler
}
