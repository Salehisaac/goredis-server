package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)


func TestNewClient1(t *testing.T){
	client,err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
		if err := client.Set(context.Background(), "foo", "1"); err != nil {
			log.Fatal(err)
		}
		val ,err := client.Get(context.Background(), "foo") 
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)
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

func TestNewClients(t *testing.T){
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(it int){
			client,err := New("localhost:5001")
			if err != nil {
				log.Fatal(err)
			}
			defer client.Close()
			key := fmt.Sprintf("client_foo_%d", it)
			value := fmt.Sprintf("client_bar_%d", it)
			if err := client.Set(context.Background(), key, value); err != nil {
				log.Fatal(err)
			}
			val ,err := client.Get(context.Background(), key) 
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("client %d got this value back => %s\n", it, val)
			wg.Done()
		}(i)

	}
	wg.Wait()
}

func (c *Client) Close()error{
	return c.conn.Close()
}