package main

import (
	"context"
	"fmt"
	"goredis/client"
	"log"
	"sync"
	"testing"
	"time"
)

func TestServerWithMultiClients(t *testing.T){
	server := NewServer(Config{})
	go func ()  {
		log.Fatal(server.Start())
	}()
	time.Sleep(time.Second)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(it int){
			c,err := client.New("localhost:5001")
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()
			key := fmt.Sprintf("client_foo_%d", it)
			value := fmt.Sprintf("client_bar_%d", it)
			if err := c.Set(context.Background(), key, value); err != nil {
				log.Fatal(err)
			}
			val ,err := c.Get(context.Background(), key) 
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("client %d got this value back => %s\n", it, val)
			wg.Done()
		}(i)

	}
	wg.Wait()

	time.Sleep(time.Second)
	if len(server.peers) != 0 {
		t.Fatalf("expexted 0 peers but got %d", len(server.peers))
	}
}

func TestRedis1(t *testing.T){
	in := map[string]string{
		"first" : "1",
		"second" : "2",
	}
	out := respWriteMap(in)
	fmt.Println(out)
}