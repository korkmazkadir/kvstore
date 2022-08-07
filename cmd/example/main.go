package main

import (
	"fmt"

	"github.com/korkmazkadir/kvstore"
)

type IDProvider struct {
	CurrentID int
}

func (i *IDProvider) NewID() int {
	c := i.CurrentID
	i.CurrentID++
	return c
}

func main() {

	spaceClient, err := kvstore.NewStoreClient("127.0.0.1", 1234)
	if err != nil {
		panic(err)
	}

	// try to register an IDProvider instance
	key := "id-provider"
	spaceClient.Put(key, IDProvider{})

	idProvider := &IDProvider{}
	// takes the IDProvider from store
	err = spaceClient.Take(key, idProvider)
	if err != nil {
		panic(err)
	}

	// gets an ID
	id := idProvider.NewID()

	fmt.Printf("Assigned ID for the process is %d\n", id)

	// put back the IDProvider so that other processes can use it
	err = spaceClient.Put(key, idProvider)
	if err != nil {
		panic(err)
	}

}
