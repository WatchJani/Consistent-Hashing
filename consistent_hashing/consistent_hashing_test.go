package consistent_hashing

import (
	"hash/crc32"
	"log"
	"testing"
)

// 73ns
func BenchmarkFindServer(b *testing.B) {
	b.StopTimer()
	ch := NewConsistentHashing()
	if err := ch.Load("./config.json"); err != nil {
		log.Println(err)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ch.FindServer("key")
	}
}

func BenchmarkSearch(b *testing.B) {
	b.StopTimer()

	ch := NewConsistentHashing()
	if err := ch.Load("./config.json"); err != nil {
		log.Println(err)
	}

	newHash := crc32.ChecksumIEEE([]byte("key"))

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Search(ch.VirtualServers, newHash)
	}
}

// 28ns
func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crc32.ChecksumIEEE([]byte("key"))
	}
}

func BenchmarkInsert(b *testing.B) {
	b.StopTimer()
	ch := NewConsistentHashing()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ch.Load("./config.json")
	}
}
