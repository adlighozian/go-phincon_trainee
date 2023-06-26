package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type rdb struct {
	redisClient *redis.Client
}

func (r *rdb) InitiateRedis() {
	r.redisClient = redis.NewClient(&redis.Options{
		// localhost:6379
		Addr:       "localhost:6379",
		ClientName: "",
		Username:   "",
		Password:   "",
		DB:         0,
		// Dial timeout for establishing new connections.
		DialTimeout: 5 * time.Second,
	})
	if r.redisClient != nil {
		fmt.Println("Redis initialization")
	}
}

func main() {
	var redisTest = rdb{}
	redisTest.InitiateRedis()

	// set string data to redis
	redisTest.RedisSetKey("username", "member_01")
	redisTest.RedisSetKey("username", "member_02")
	redisTest.RedisSetKey("password", "Test123")

	// set integer data to redis
	num := 12313214
	str := strconv.Itoa(num)
	redisTest.RedisSetKey("token", str)

	// set array data to redis
	someArray := [3]string{"member_01", "member_02", "member_03"}
	byteArray, _ := json.Marshal(someArray)
	redisTest.RedisSetKey("user_array", byteArray)

	// set slice data to redis
	someSlice := []string{"member_01", "member_02", "member_03"}
	byteSlice, _ := json.Marshal(someSlice)
	redisTest.RedisSetKey("user_slice", byteSlice)

	// set map data to redis
	user := map[string]interface{}{
		"username": "member_01", "password": "Tes123", "token": 1214113141,
	}
	byteMap, _ := json.Marshal(user)
	redisTest.RedisSetKey("user_map", byteMap)

	// set struct data to redis
	type userdata struct {
		Username string
		Password string
		Token    int
	}
	user2 := userdata{"member_01", "Test123", 43234312432413}
	byteStruct, _ := json.Marshal(user2)
	redisTest.RedisSetKey("user_struct", byteStruct)

	// get value by key
	keys := []string{"username", "password", "token", "user_array", "user_slice", "user_map", "user_struct"}
	for _, key := range keys {
		value, err := redisTest.RedisGetKey(key)
		if err != nil {
			fmt.Printf("Error get value of key %s \n", value)
			continue
		}
		fmt.Printf("Value of username key %s is %s\n", key, value)
	}

	// delete value by key
	redisTest.RedisDeleteKey("username")
	redisTest.RedisDeleteKey("password")
}

func (r rdb) RedisSetKey(key string, value interface{}) {
	err := r.redisClient.Set(ctx, key, value, 5*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Key %s with value %v is set\n", key, value)
}

func (r rdb) RedisGetKey(key string) (value string, err error) {
	value, err = r.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		err = fmt.Errorf("key doesn't exist : %s, data type : %T", redis.Nil, redis.Nil)
		return
	} else if err != nil {
		return
	} else {
		return
	}
}

func (r rdb) RedisDeleteKey(key string) {
	if err := r.redisClient.Get(ctx, key).Err(); err != nil {
		if err = r.redisClient.Del(ctx, key).Err(); err != nil {
			fmt.Printf("Error delete : %s\n", err.Error())
		}
		fmt.Printf("Success delete key %s\n", key)
	}
}
