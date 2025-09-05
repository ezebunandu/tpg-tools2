package kv

type store struct {
    data map[string]string
}

func Openstore(path string) (*store, error){
    return &store{data: map[string]string{}}, nil
}

func (s *store) Get(key string) (string, bool){
    value, ok := s.data[key]
    return value, ok
}

func (s *store) Set(key, value string){
    s.data[key] = value
}
