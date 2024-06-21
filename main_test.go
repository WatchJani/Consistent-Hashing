package main

// import "testing"

// func BenchmarkInsert(b *testing.B) {
// 	b.StopTimer()

// 	numberOfServer := 50
// 	virtualServer := numberOfServer * VNServer

// 	p := make([]Node, virtualServer)

// 	b.StartTimer()

// 	for i := 0; i < b.N; i++ {
// 		LoadServer(p)
// 	}
// }

// // 18.46ns
// func BenchmarkSearch(b *testing.B) {
// 	b.StopTimer()

// 	numberOfServer := 50
// 	virtualServer := numberOfServer * VNServer

// 	p := make([]Node, virtualServer)
// 	LoadServer(p)

// 	b.StartTimer()

// 	for i := 0; i < b.N; i++ {
// 		BinarySearchFirstGreaterEqual(p, 1056115201)
// 	}
// }
