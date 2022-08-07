package kvstore

import "fmt"

type ErrKeyNotExist struct {
	Key string
}

func (e *ErrKeyNotExist) Error() string {

	return fmt.Sprintf("key %s does not exist", e.Key)
}

type ErrKeyAlreadyExist struct {
	Key string
}

func (e *ErrKeyAlreadyExist) Error() string {

	return fmt.Sprintf("key %s already exists", e.Key)
}
