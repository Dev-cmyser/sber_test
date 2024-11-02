package cache

import (
	"time"
)

func SetCache(ttl, amountKeys int) {

	cache := expirable.NewLRU[string, string](5, nil, time.Millisecond*10)

}
