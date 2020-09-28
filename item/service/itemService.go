package service

import (
	"fmt"

	"github.com/balsuvendukumar/item/domain"
	"github.com/balsuvendukumar/item/util"
)

func InsertItem(r domain.ItemDetail) (*domain.ItemDetail, *util.ItemError) {
	return domain.InsertItem(r)

}

func GetItem(key int) (*domain.ItemDetail, *util.ItemError) {
	fmt.Println(key)
	return domain.FetchOneItem(key)

}

/*
func GetItems()(*domain.GetItem,utils.ItemError){


}
*/
