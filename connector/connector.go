package connector
import (
	"github.com/go-redis/redis"
	"golang.org/x/net/context"
	"log"
	"math/big"
)

var ctx = context.Background()

// SetBigKey add to redis fibonacci number
// key --  number position
// value -- fibonacci number value
// Values stored as a strings
func SetBigKey(key int, value big.Int) {
	rdb := redis.NewClient(&redis.Options{ //todo make global config
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,  // use default DB
	})
	err := rdb.Do(ctx, "SET", key, value.String()).Err()
	if err != nil {
		log.Printf("cant get value to redis; key %d", key)
	}
	log.Println("set Value to Cache")
}

// GetBigKey Try to get fibonacci number with position n from redis
// If there is not key in redis returns nil
func GetBigKey(key int) *big.Int {
	rdb := redis.NewClient(&redis.Options{  //todo make global config
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,  // use default DB
	})
	val, err := rdb.Do(ctx, "GET", key).Result()
	if err == redis.Nil {
		log.Printf("cant get value from redis; key %d", key)
		return nil
	}
	n := new(big.Int)
	n.SetString(val.(string), 10)
	return n
}