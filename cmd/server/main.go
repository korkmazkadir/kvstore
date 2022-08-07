package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/korkmazkadir/kvstore"
)

func main() {

	store := kvstore.NewStore()
	spaceServer := kvstore.NewServer(store)

	err := rpc.Register(spaceServer)
	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal(e)
	}

	log.Printf("service started and listening on :1234\n")

	for {
		conn, _ := l.Accept()
		go func() {
			rpc.ServeConn(conn)
		}()
	}

}
