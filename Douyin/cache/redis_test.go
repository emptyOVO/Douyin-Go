package cache

import (
	"Douyin/config"
	"fmt"
	"testing"
)

func TestSetUserCount(t *testing.T) {
	err := config.ConfInit()
	if err != nil {
		return
	}
	err = RedisPoolInit()
	if err != nil {
		return
	}
	key := IsUserRelation(1, 2)
	fmt.Println(key)
	err = IncrByUserTotalFavorite(13)

}
