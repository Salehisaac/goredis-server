package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)


func TestNewClient1(t *testing.T){
	client,err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	time.Sleep(time.Second)
		if err := client.Set(context.Background(), "foo", 1); err != nil {
			log.Fatal(err)
		}
		val ,err := client.Get(context.Background(), "foo") 
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)
}


func TestRedis(t *testing.T){
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:5001",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(context.Background(), "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(context.Background(), "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)
}


func TestNewClient(t *testing.T){
	client,err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		
		
		if err := client.Set(context.Background(), fmt.Sprintf("foo_%d", i), fmt.Sprintf("bar_%d", i)); err != nil {
			log.Fatal(err)
		}
		fmt.Println("SET => " , fmt.Sprintf("bar_%d", i))
		val ,err := client.Get(context.Background(), fmt.Sprintf("foo_%d", i)) 
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("GET => " , val)
	}
}

