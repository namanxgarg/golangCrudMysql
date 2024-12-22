package main

import (
	"golangCrudMysql/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}