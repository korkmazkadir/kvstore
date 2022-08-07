package kvstore

import (
	"encoding/json"
	"fmt"
	"net/rpc"
)

type StoreClient struct {
	rpcClient *rpc.Client
}

func NewStoreClient(ipAddress string, portNumber int) (*StoreClient, error) {

	rpcClient, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", ipAddress, portNumber))
	if err != nil {
		return nil, err
	}

	spaceClient := &StoreClient{rpcClient: rpcClient}

	return spaceClient, nil
}

func (sc *StoreClient) Put(key string, v interface{}) error {

	objBytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	tuple := KVPair{Key: key, Value: string(objBytes)}

	err = sc.rpcClient.Call("StoreServer.Put", tuple, nil)
	return err
}

func (sc *StoreClient) Get(key string, v interface{}) error {

	var objString string
	err := sc.rpcClient.Call("StoreServer.Get", key, &objString)
	if err != nil {
		return err
	}

	ptr, ok := v.(*string)
	if ok {
		*ptr = objString
		return nil
	}

	err = json.Unmarshal([]byte(objString), v)
	return err
}

func (sc *StoreClient) Copy(fromKey string, toKey string) error {

	copyReq := CopyMoveRequest{FromKey: fromKey, ToKey: toKey}
	err := sc.rpcClient.Call("StoreServer.Copy", copyReq, nil)
	return err
}

func (sc *StoreClient) Move(fromKey string, toKey string) error {

	copyReq := CopyMoveRequest{FromKey: fromKey, ToKey: toKey}
	err := sc.rpcClient.Call("StoreServer.Move", copyReq, nil)
	return err
}

func (sc *StoreClient) Take(key string, v interface{}) error {

	var objString string
	err := sc.rpcClient.Call("StoreServer.Take", key, &objString)
	if err != nil {
		return err
	}

	ptr, ok := v.(*string)
	if ok {
		*ptr = objString
		return nil
	}

	err = json.Unmarshal([]byte(objString), v)
	return err
}

func (sc *StoreClient) List(regExp string) ([]string, error) {

	var keyList []string
	err := sc.rpcClient.Call("StoreServer.List", regExp, &keyList)
	if err != nil {
		return nil, err
	}

	return keyList, nil
}
