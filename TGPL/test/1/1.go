package main

import "fmt"

func del(a []string) []string {
	i := 0
	for _, val := range a {
		if val != "" {
			a[i] = val
			i++
		}
	}
	return a[:i]
}

func pardon(a []string) []string {
	res := make([]string, 0)
	status := 1
	for _, val := range a {
		for _, val2 := range res {
			if val2 == val {
				status = 0
				break
			}
		}
		if status == 1 {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	x := []string{"red", "", "blue", "", "", "greee", "blue", "green", "red", "red", "red"}
	x = del(x)
	fmt.Println(x)
	x = pardon(x)
	fmt.Println(x)
}
