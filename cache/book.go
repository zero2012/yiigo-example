package cache

import (
	"demo/models"
	"encoding/json"

	"github.com/iiinsomnia/yiigo"
)

// GetBookCache 获取缓存
func GetBookCache(id int, data *models.Book) bool {
	conn := yiigo.Redis.Get()
	defer conn.Close()

	r, err := conn.Do("HGET", "slim:books", id)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	if r == nil {
		return false
	}

	err = yiigo.ScanJSON(r, data)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}

// SetBookCache 设置缓存
func SetBookCache(id int, data *models.Book) bool {
	conn := yiigo.Redis.Get()
	defer conn.Close()

	cache, err := json.Marshal(data)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	_, err = conn.Do("HSET", "slim:books", id, cache)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}

// DelBookCache 删除缓存
func DelBookCache(id int) bool {
	conn := yiigo.Redis.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", "slim:books", id)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}
