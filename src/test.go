package main

import "fmt"

func Square(a1, a2 []int) int {
	x := a2[0] - a1[0]
	y := a2[1] - a1[1]
	return x*x + y*y
}

func numberOfBoomerangs(points [][]int) int {

	res := 0
	size := len(points)
	if size < 3 {
		return 0
	}

	for i := 0; i < size; i++ {
		dmap := make(map[int]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}
			d := Square(points[i], points[j])
			if _, ok := dmap[d]; ok {
				dmap[d]++
			} else {
				dmap[d] = 1
			}
		}
		for _, v := range dmap {
			res += v * (v - 1)
		}
	}
	return res
}

func main() {
	var a [][]int
	a = [][]int{{2, 2}, {2, 1}}
	fmt.Println(numberOfBoomerangs(a))

	// fmt.Println(a[1])
}
