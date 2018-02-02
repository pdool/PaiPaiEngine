package DataStruct

// Set ...
type Set struct {
	items map[interface{}]struct{}
}

// Add ...
func (s *Set) Add(element interface{}) bool {
	if !s.Has(element) {
		if s.items == nil {
			s.items = make(map[interface{}]struct{})
		}
		s.items[element] = struct{}{}
		return true
	}
	return false
}

// Remove ...
func (s *Set) Remove(element interface{}) bool {
	if s.Has(element) {
		delete(s.items, element)
		return true
	}
	return false
}

// Has ...
func (s *Set) Has(element interface{}) bool {
	_, has := s.items[element]
	return has
}

// Clear ...
func (s *Set) Clear() {
	s.items = make(map[interface{}]struct{})
}

// Size ...
func (s *Set) Size() int {
	if s.items == nil {
		return 0
	}
	cnt := 0
	for range s.items {
		cnt++
	}
	return cnt
}

// Values ...
func (s *Set) Values() []interface{} {
	var v []interface{}
	for k := range s.items {
		v = append(v, k)
	}
	return v
}

// Union ...
func (s *Set) Union(otherSet *Set) *Set {
	newSet := Set{}
	for i, v := 0, s.Values(); i < len(v); i++ {
		newSet.Add(v[i])
	}
	for i, v := 0, otherSet.Values(); i < len(v); i++ {
		newSet.Add(v[i])
	}
	return &newSet
}

// Intersection ...
func (s *Set) Intersection(otherSet *Set) *Set {
	newSet := Set{}
	for i, v := 0, s.Values(); i < len(v); i++ {
		if otherSet.Has(v[i]) {
			newSet.Add(v[i])
		}
	}
	return &newSet
}

// Difference ...
func (s *Set) Difference(otherSet *Set) *Set {
	newSet := Set{}
	for i, v := 0, s.Values(); i < len(v); i++ {
		if !otherSet.Has(v[i]) {
			newSet.Add(v[i])
		}
	}
	return &newSet
}

// Subset ...
func (s *Set) Subset(otherSet *Set) bool {
	if s.Size() > otherSet.Size() {
		return false
	}
	for i, v := 0, s.Values(); i < len(v); i++ {
		if !otherSet.Has(v[i]) {
			return false
		}
	}
	return true
}
