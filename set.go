package sets

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](size int) Set[T] {
	return make(Set[T], size)
}

func (s Set[T]) Add(key T) {
	s[key] = struct{}{}
}

func (s Set[T]) Remove(key T) {
	delete(s, key)
}

func (s Set[T]) Contains(key T) bool {
	_, has := s[key]
	return has
}

func (s Set[T]) Values() (values []T) {
	for v := range s {
		values = append(values, v)
	}

	return
}
