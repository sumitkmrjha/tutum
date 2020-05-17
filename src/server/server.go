package server

import (
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/sumitkmrjha/tutum/src/core"
)

func GetDoc(w http.ResponseWriter, req *http.Request){
	//params := mux.Vars(req)
	http.Redirect(w, req, "" , 301)
}

func CreateDoc(w http.ResponseWriter, req *http.Request){
	//params := mux.Vars(req)
	//res := core.insert(params["key"], params["value"])
	http.Redirect(w, req, "", 301)
}

func initRoutes(){
	router := mux.NewRouter()
	router.HandleFunc("/{id}",GetDoc).Methods("GET")
	router.HandleFunc("create", CreateDoc).Methods("PUT")
}


