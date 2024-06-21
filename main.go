package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Node struct {
	Num int
}

type ConsistentHashing struct {
	NumberVirtualNode int
}

const VNServer = 256

func main() {
	numberOfServer := 1
	virtualServer := numberOfServer * VNServer

	p := make([]Node, virtualServer)

	LoadServer(p)

	for index, value := range p {
		fmt.Println(index, value.Num)
	}

	fmt.Println("my result", BinarySearchFirstGreaterEqual(p, 1056115201))
}

func LoadServer(p []Node) {
	for i := range p {
		p[i] = Node{
			Num: rand.Intn(2147483648),
		}
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].Num < p[j].Num
	})
}

func BinarySearchFirstGreaterEqual(arr []Node, target int) int {
	left, right := 0, len(arr)-1
	result := len(arr)

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid].Num >= target {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}
