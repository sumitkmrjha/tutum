package server

import (
	"github.com/gorilla/mux"
	"github.com/sumitkmrjha/tutum/src/core"
	"net/http"
)


func GetDoc(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	t := core.GetTxn()
	doc := t.GetDoc([]byte(params["key"]),nil)
	http.Redirect(w, req, doc.String(), 301)
}

func CreateDoc(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	t := core.GetTxn()
	err := t.InsertDoc([]byte(params["key"]),[]byte(params["value"]), nil)
	if(err != nil) {
		http.Redirect(w, req, "", 301)
	}else{
		http.Redirect(w, req, err.Error(),  420)
	}
}

func UpdateDoc(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	t := core.GetTxn()
	err := t.UpdateDoc([]byte(params["key"]),[]byte(params["value"]), nil)
	if(err != nil) {
		http.Redirect(w, req, "", 301)
	}else{
		http.Redirect(w, req, err.Error(),  420)
	}

}

func DeleteDoc(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	t := core.GetTxn()
	err := t.DeleteDoc([]byte(params["key"]), nil)
	if(err != nil) {
		http.Redirect(w, req, "", 301)
	}else{
		http.Redirect(w, req, err.Error(),  420)
	}
}

func initRoutes(){
	router := mux.NewRouter()

	/******************** Doc APIs *******************************/
	router.HandleFunc("/{id}",GetDoc).Methods("GET")
	router.HandleFunc("create", CreateDoc).Methods("PUT")
	router.HandleFunc("update", UpdateDoc).Methods("PUT")
	router.HandleFunc("delete", DeleteDoc).Methods("PUT")


	/********************* Collection APIs **********************/




}


