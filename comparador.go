package main

func comparator(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := makeMap(a)
	n := makeMap(b)

	for key, _ := range m {
		if m[key] != n[key] {
			return false
		}
	}
	return true
}

func makeMap(word string) map[string]int {
	x := make(map[string]int)
	for i := 0; i < len(word); i++ {
		letra := string(word[i])
		_, exists := x[letra]
		if exists {
			x[word]++
		} else {
			x[word] = 1
		}
	}
	return x
}
