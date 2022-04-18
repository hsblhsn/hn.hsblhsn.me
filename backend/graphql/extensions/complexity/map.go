package complexity

type Map map[string]int

func (m Map) Get(key string, def int) int {
	val, ok := m[key]
	if !ok {
		return def
	}
	return val
}
