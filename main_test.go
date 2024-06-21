package main

import "testing"

func Benchmark(b *testing.B) {
	b.StopTimer()

	numberOfServer := 50
	virtualServer := numberOfServer * VNServer

	p := make([]Node, virtualServer)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		LoadServer(p)
	}
}
