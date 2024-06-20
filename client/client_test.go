package client

import(
	"testing"
	"log"
	"time"
	"context"
	"fmt"
)


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