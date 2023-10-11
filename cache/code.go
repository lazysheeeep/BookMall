package cache

import (
	"context"
	"strconv"
	"time"
)

func AddSmsCode(ctx context.Context, uId uint, phoneNum string, code string) error { //添加验证码code到redis数据库，设置有效时长
	iId := strconv.Itoa(int(uId))
	key := "Sms:user:" + iId + ":" + phoneNum
	_, err := RedisClient.SetEx(ctx, key, code, 2*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetSmsCode(ctx context.Context, uId uint, phoneNum string) string {
	iId := strconv.Itoa(int(uId))
	key := "Sms:user:" + iId + ":" + phoneNum
	result, err := RedisClient.Get(ctx, key).Result() //It returns redis.Nil error when key does not exist.
	if err != nil {
		return ""
	}
	return result
}
