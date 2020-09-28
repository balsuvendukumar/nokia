package app

import (
	"net/http"

	"github.com/balsuvendukumar/item/controller"
	"github.com/gorilla/mux"
)

func StartApp() {
	router := mux.NewRouter()
	router.HandleFunc("/getOneItem/{item_id}", controller.GetItemDetail).Methods("GET")
	router.HandleFunc("/createItem", controller.CreateItem).Methods("POST")
	//router.HandleFunc("/getAllItems", controller.GetItemDetail).Methods("GET")
	http.ListenAndServe(":8080", router)

}
