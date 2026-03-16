package uniquer

func Unique(i []int) []int {
	seen := make(map[int]bool)
	var store []int
	for _, value := range i {
		if _, ok := seen[value]; !ok {
			seen[value] = true
			store = append(store, value)
		}
	}
	return store
}
