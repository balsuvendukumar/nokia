package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/balsuvendukumar/item/cache"
	"github.com/balsuvendukumar/item/domain"
	"github.com/balsuvendukumar/item/service"
	"github.com/balsuvendukumar/item/util"
	"github.com/gorilla/mux"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var requestCame domain.ItemDetail
	err := json.NewDecoder(r.Body).Decode(&requestCame)
	if err != nil {
		er := &util.ItemError{
			http.StatusBadRequest, "Error-Bad Json Request",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(er)
		return

	}
	detail, errorRes := service.InsertItem(requestCame)
	if errorRes != nil {
		json.NewEncoder(w).Encode(errorRes)
		return

	}

	cKey := strconv.Itoa(int(detail.ItemID))
	cache.SetCache(cKey, detail)
	fmt.Println("setting Cache suucessful")
	json.NewEncoder(w).Encode(detail)

}

func GetItemDetail(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)
	key := name["item_id"]

	oneItemDetail, found := cache.GetCache(key)
	if found {
		fmt.Println("reading from cache")
		json.NewEncoder(w).Encode(oneItemDetail)
		return
	}

	keyIsInt, err := strconv.Atoi(key)
	if err != nil {
		er := &util.ItemError{
			http.StatusBadRequest, "Key in request is not numeric",
		}
		json.NewEncoder(w).Encode(er)
		return
	}
	oneItemDetail, er := service.GetItem(int(keyIsInt))

	if er != nil {
		json.NewEncoder(w).Encode(er)
		return
	}

	json.NewEncoder(w).Encode(oneItemDetail)

}
