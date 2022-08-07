package kvstore

import (
	"log"
	"regexp"
	"sync"
)

type Store struct {
	// protects following fields
	mutex sync.Mutex
	// objects are stored in a map
	objectMap map[string]string
}

func NewStore() *Store {
	return &Store{objectMap: make(map[string]string)}
}

// Private methods

func (s *Store) put(key string, value string) error {

	_, ok := s.objectMap[key]
	if ok {
		// key already exists, throws an error
		return &ErrKeyAlreadyExist{Key: key}
	}

	s.objectMap[key] = value

	return nil
}

func (s *Store) get(key string) (string, error) {

	obj, ok := s.objectMap[key]
	if ok == false {
		// key does not exist, throws an error
		return "", &ErrKeyNotExist{Key: key}
	}

	return obj, nil
}

// Public methods

func (s *Store) Put(key string, value string) error {

	// protects against concurrent access
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("put %s", key)

	return s.put(key, value)
}

func (s *Store) Get(key string) (string, error) {

	// protects against concurrent access
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("get %s", key)

	return s.get(key)
}

func (s *Store) Take(key string) (string, error) {

	// protects against concurrent access
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("take %s", key)

	obj, err := s.get(key)
	if err != nil {
		return "", err
	}

	// removes the key from the space
	delete(s.objectMap, key)

	return obj, nil
}

func (s *Store) Copy(fromKey string, toKey string) error {

	// protects against concurrent access
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("copy from %s to %s", fromKey, toKey)

	obj, err := s.get(fromKey)
	if err != nil {
		return err
	}

	err = s.put(toKey, obj)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Move(fromKey string, toKey string) error {

	// protects against concurrent access
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("move from %s to %s", fromKey, toKey)

	obj, err := s.get(fromKey)
	if err != nil {
		return err
	}

	err = s.put(toKey, obj)
	if err != nil {
		return err
	}

	// removes the key from the store
	delete(s.objectMap, fromKey)

	return nil
}

func (s *Store) List(regExp string) ([]string, error) {

	// protects against concurrent access
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("list regexp %s", regExp)

	var r *regexp.Regexp
	if regExp != "" {
		var err error
		r, err = regexp.Compile(regExp)
		if err != nil {
			return nil, err
		}
	}

	var result []string
	for key := range s.objectMap {
		if r != nil {
			if r.MatchString(key) {
				result = append(result, key)
			}
		} else {
			result = append(result, key)
		}
	}

	return result, nil
}
