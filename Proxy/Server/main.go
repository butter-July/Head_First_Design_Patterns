package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

func main() {
	machine := NewGumballMachine(2) // 正确调用
	remote := &GumballRemote{machine: machine}

	rpc.Register(remote)
	rpc.HandleHTTP()

	fmt.Println("Starting RPC server on port 8080...")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println("Error starting RPC server:", err)
	}

}
