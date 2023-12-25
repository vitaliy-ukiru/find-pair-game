package container

type Hashset[K comparable] map[K]struct{}

func HashSetFromSlice[K comparable](items []K) Hashset[K] {
	h := make(Hashset[K])
	for _, item := range items {
		h.Add(item)
	}
	return h
}

func (h Hashset[K]) Add(item K) {
	h[item] = struct{}{}
}
func (h Hashset[K]) Has(item K) bool {
	_, ok := h[item]
	return ok
}

func (h Hashset[K]) Delete(item K) {
	delete(h, item)
}

func (h Hashset[K]) ToSlice() []K {
	slice := make([]K, 0, len(h))
	for k := range h {
		slice = append(slice, k)
	}
	return slice
}
