package tools

// return set (same array but with unique keys)
func Set[A int | string](arr []A) []A {
	var myMap = make(map[A]bool)
	var newSet []A

	for _, v := range arr {
		if _, ok := myMap[v]; !ok {
			myMap[v] = true
		}
	}

	for k := range myMap {
		newSet = append(newSet, k)
	}

	return newSet
}
