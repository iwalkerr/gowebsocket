package cache

import (
	"fmt"
	"gowebsocket/lib/redislib"
)

// 黑名单
const (
	backlistPrefix = "acc:user:backlist" // 数据不重复提交
)

func getBacklistKey() (key string) {
	key = fmt.Sprintf("%s", backlistPrefix)

	return
}

// 设置黑名单用户
func SetUserBacklist(uid int) error {
	key := getBacklistKey()

	redisClient := redislib.GetClient()
	_, err := redisClient.Do("hSet", key, uid, "1").Int()
	return err
}

func CheckBacklist(uid int) (bool, error) {
	key := getBacklistKey()
	redisClient := redislib.GetClient()
	number, err := redisClient.Do("hExists", key, uid).Int()
	if number == 1 {
		return true, err
	}
	return false, err

}
