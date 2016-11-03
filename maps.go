package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 83
	m["k2"] = 78347

	fmt.Println("map:", m)
	fmt.Println("maplen:", len(m))

	v := m["k1"]
	fmt.Println("v:", v)

	delete(m, "k2")
	fmt.Println("mapnew:", m)

	v1, stat1 := m["k1"]
	fmt.Println("Check:", v1, ",", stat1)
	v2, stat2 := m["k2"]
	fmt.Println("Check:", v2, ",", stat2)

	d := map[string]int{"key1": 3, "key2": 4}
	fmt.Println("dyn:", d)
}
