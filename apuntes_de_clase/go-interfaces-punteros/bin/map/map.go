package _map

func main() {
	m := map[string]any{
		"Int":   1,
		"Float": 3.14,
		"Array": []any{1, 2, 3, 4, 5},
	}

	value, ok := m["Bool"]
	if !ok {
		println("Key not found")
		return
	}

	f, ok := value.(float64)
	if !ok {
		println("Value is not a float64")
		return
	}
	result := f + 2
	println(result)
}
