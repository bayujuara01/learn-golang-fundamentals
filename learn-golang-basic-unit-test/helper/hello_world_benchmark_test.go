package helper

import (
	"testing"
)

/**
Run :
- go test ./... -run=TestYangTidakAda -bench=.
- go test ./... -run=TestYangTidakAda -bench=BenchmarkHello
*/

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bayu")
	}
}

func BenchmarkHelloWorldBayu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bayu")
	}
}

func BenchmarkHelloWorldSubBench(b *testing.B) {
	b.Run("Sub Benchmark 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bayu")
		}
	})

	b.Run("Sub Benchmark 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bayu")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchTables := []struct {
		name    string
		request string
	}{
		{name: "Bench_1", request: "Bayu_1"},
		{name: "Bench_2", request: "Bayu_2"},
		{name: "Bench_3", request: "Bayu_3"},
	}

	for _, bench := range benchTables {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.request)
			}
		})
	}
}
