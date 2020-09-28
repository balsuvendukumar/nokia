package cache

import (
	"time"

	"github.com/balsuvendukumar/item/domain"
	"github.com/patrickmn/go-cache"
)

var cach = cache.New(5*time.Minute, 5*time.Minute)

func SetCache(key string, detail *domain.ItemDetail) {
	cach.Set(key, detail, cache.NoExpiration)
}

func GetCache(key string) (*domain.ItemDetail, bool) {
	var detail *domain.ItemDetail
	var found bool
	data, found := cach.Get(key)
	if found {
		detail = data.(*domain.ItemDetail)
	}
	return detail, found
}
