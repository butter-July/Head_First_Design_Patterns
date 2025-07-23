package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal("Dialing:", err)
	}
	args := struct{}{}
	var reply string
	err = client.Call("GumballRemote.InsertQuarter", &args, &reply)
	if err != nil {
		log.Fatal("InsertQuarter error:", err)
	}
	fmt.Println(reply)

}
