package set

type Set[Item comparable] map[Item]struct{}

func (s Set[Item]) Add(items ...Item) Set[Item] {
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func (s Set[Item]) Has(item Item) bool {
	_, has := s[item]
	return has
}

func (s Set[Item]) Remove(items ...Item) Set[Item] {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Sub 返回一个新的Set,是原来的Set中去除集合b中的元素
func (s Set[Item]) Sub(b Set[Item]) Set[Item] {
	res := NewSet[Item](0)
	for item := range s {
		if !b.Has(item) {
			res.Add(item)
		}
	}
	return res
}

// IsSubset 判断当前集合是否是另一个集合的子集
func (s Set[Item]) IsSubset(b Set[Item]) bool {
	for item := range s {
		if !b.Has(item) {
			return false
		}
	}
	return true
}

func (s Set[Item]) HasAll(items ...Item) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

func (s Set[Item]) Intersect(b Set[Item]) Set[Item] {
	res := NewSet[Item](0)
	for item := range s {
		if b.Has(item) {
			res.Add(item)
		}
	}
	return res
}

func (s Set[Item]) Union(b Set[Item]) Set[Item] {
	res := NewSet[Item](0)
	for item := range s {
		res.Add(item)
	}
	for item := range b {
		res.Add(item)
	}
	return res
}

func (s Set[Item]) Count() int {
	return len(s)
}

func (s Set[Item]) ToSlice() []Item {
	items := make([]Item, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}

// NewSet cap<=0零时，使用默认的map容量
func NewSet[Key comparable](cap int) Set[Key] {
	if cap > 0 {
		return make(Set[Key], cap)
	}
	return make(Set[Key])
}

func NewSetFromItems[Key comparable](items ...Key) Set[Key] {
	s := NewSet[Key](len(items))
	for _, item := range items {
		s.Add(item)
	}
	return s
}
