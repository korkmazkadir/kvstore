package kvstore

type StoreServer struct {
	space *Store
}

type KVPair struct {
	Key   string
	Value string
}

type CopyMoveRequest struct {
	FromKey string
	ToKey   string
}

func NewServer(space *Store) *StoreServer {

	server := &StoreServer{space: space}
	return server
}

func (s *StoreServer) Put(t KVPair, reply *string) error {

	err := s.space.Put(t.Key, t.Value)
	return err
}

func (s *StoreServer) Get(key string, reply *string) error {

	obj, err := s.space.Get(key)
	*reply = obj
	return err
}

func (s *StoreServer) Copy(copyReq CopyMoveRequest, reply *string) error {

	err := s.space.Copy(copyReq.FromKey, copyReq.ToKey)
	return err
}

func (s *StoreServer) Move(moveReq CopyMoveRequest, reply *string) error {

	err := s.space.Move(moveReq.FromKey, moveReq.ToKey)
	return err
}

func (s *StoreServer) Take(key string, reply *string) error {

	obj, err := s.space.Take(key)
	*reply = obj
	return err
}

func (s *StoreServer) List(regExp string, reply *[]string) error {

	keyList, err := s.space.List(regExp)
	*reply = keyList
	return err
}
