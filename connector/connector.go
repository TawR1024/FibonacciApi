package connector

import (
	"github.com/go-redis/redis"
	"golang.org/x/net/context"
	"log"
	"math/big"
	"strconv"
)

var ctx = context.Background()

// SetBigKey add to redis fibonacci number
// key --  number position
// value -- fibonacci number value
// Values stored as a strings
func SetBigKey(rdb *redis.Client, key int, value big.Int) {
	err := rdb.Set(ctx, strconv.Itoa(key), value.String(), 0).Err()
	if err != nil {
		panic(err)
	}
}

// GetBigKey Try to get fibonacci number with position n from redis
// If there is not key in redis returns nil
func GetBigKey(rdb *redis.Client, key int) *big.Int {
	//rdb := redis.NewClient(&redis.Options{  //todo make global config
	//	Addr:     "localhost:6379",
	//	Password: "",
	//	DB:       0,  // use default DB
	//})
	val, err := rdb.Get(ctx, strconv.Itoa(key)).Result()
	if err != nil || val == "" {
		log.Printf("cant get value from redis; key %d, value is %v", key, val)
		return nil
	}
	n := new(big.Int)
	n.SetString(val, 10)
	return n
}
