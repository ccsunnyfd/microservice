package data

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"log"
	"microservice/api/user/service/v1"
	"testing"
	"time"
)

var rdb *redis.Client

func setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         "172.16.20.198:31971",
		Password:     "",
		DB:           0,
		DialTimeout:  1 * time.Second,
		WriteTimeout: 600 * time.Millisecond,
		ReadTimeout:  400 * time.Millisecond,
	})
	rdb.AddHook(redisotel.TracingHook{})
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	log.Println(pong, err)
}

func TestUserRepo_SetUserInfo(t *testing.T) {
	setup()
	ctx := context.Background()
	data, _ := json.Marshal(&v1.UserInfo{
		ID:       993,
		Username: "jfkj",
		Password: "jf",
		RealName: "ff",
		Mobile:   "34894893",
		Email:    "fdf@fdf.com",
	})
	err := rdb.Set(ctx, "token", data , 6 * time.Second).Err()
	if err != nil {
		log.Printf("redis set error: %v\n", err)
	}
	log.Printf("redis set success.\n")

	get, err := rdb.Get(ctx, "token").Bytes()
	if err != nil {
		log.Printf("redis get error: %v\n", err)
	}
	userInfoReply := &v1.UserInfo{}
	err = json.Unmarshal(get, userInfoReply)
	if err != nil {
		log.Printf("redis get unmarshal error: %v\n", err)
	} else {
		log.Printf("redis get: %v\n", userInfoReply)
	}

	log.Println("message", "closing the data resources")
	if err := rdb.Close(); err != nil {
		log.Printf("close redis conn error: %v\n", err)
	}
}