package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/korkmazkadir/kvstore"
)

func main() {

	cmd := flag.String("command", "", "command")
	key := flag.String("key", "", "key")
	fromKey := flag.String("from", "", "from key")
	toKey := flag.String("to", "", "to key")
	value := flag.String("value", "", "value")
	regExp := flag.String("regexp", "", "regular expression to filter the list")

	flag.Parse()

	switch *cmd {
	case "put":
		if *key == "" || *value == "" {
			fmt.Println("key and value parameters can not be empty for put operation")
			return
		}

		put(*key, *value)

	case "get":
		if *key == "" {
			fmt.Println("key parameter can not be empty for get operation")
			return
		}

		get(*key)

	case "copy":
		if *fromKey == "" || *toKey == "" {
			fmt.Println("fromKey and toKey parameters can not be empty for copy operation")
			return
		}

		copy(*fromKey, *toKey)

	case "move":
		if *fromKey == "" || *toKey == "" {
			fmt.Println("fromKey and toKey parameters can not be empty for move operation")
			return
		}

		move(*fromKey, *toKey)

	case "take":
		if *key == "" {
			fmt.Println("for take command key parameter can not be empty")
			return
		}

		take(*key)

	case "list":

		list(*regExp)

	default:
		fmt.Printf("unknown command %s\n", *cmd)
	}

}

func client() *kvstore.StoreClient {
	spaceClient, err := kvstore.NewStoreClient("127.0.0.1", 1234)
	if err != nil {
		panic(err)
	}
	return spaceClient
}

func put(key string, value string) {

	spaceClient := client()

	err := spaceClient.Put(key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
}

func get(key string) {
	spaceClient := client()

	var result string
	err := spaceClient.Get(key, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", result)
}

func copy(fromKey string, toKey string) {
	spaceClient := client()

	err := spaceClient.Copy(fromKey, toKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
}

func move(fromKey string, toKey string) {
	spaceClient := client()

	err := spaceClient.Move(fromKey, toKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
}

func take(key string) {
	spaceClient := client()

	var result string
	err := spaceClient.Take(key, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", result)
}

func list(regExp string) {
	spaceClient := client()

	keyList, err := spaceClient.List(regExp)
	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Strings(keyList)
	for i, k := range keyList {
		fmt.Printf("[%d]%s\n", i+1, k)
	}
}
