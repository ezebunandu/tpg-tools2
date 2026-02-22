package kv

type store struct {
    data map[string]string
}

func OpenStore(path string) (*store, error) {
	return &store{
        data: map[string]string{},
    }, nil
}

func (s *store) Get(key string) (string, bool) {
    v, ok := s.data[key]
	return v, ok
}

func (s *store) Set(k, v string) {
    s.data[k] = v
}
