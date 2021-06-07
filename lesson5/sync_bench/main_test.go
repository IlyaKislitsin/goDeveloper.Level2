package main

import (
	"math/rand"
	"testing"
)

func BenchmarkSetAdd10Has90(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setAddOrHas(set, 0.1)
			}
		})
	})
}

func BenchmarkSetAdd50Has50(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setAddOrHas(set, 0.5)
			}
		})
	})
}

func BenchmarkSetAdd90Has10(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setAddOrHas(set, 0.9)
			}
		})
	})
}

func BenchmarkRWSetAdd10Has90(b *testing.B) {
	var set = NewRWSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setAddOrHas(set, 0.1)
			}
		})
	})
}

func BenchmarkRWSetAdd50Has50(b *testing.B) {
	var set = NewRWSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setAddOrHas(set, 0.5)
			}
		})
	})
}

func BenchmarkRWSetAdd90Has10(b *testing.B) {
	var set = NewRWSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setAddOrHas(set, 0.9)
			}
		})
	})
}

func BenchmarkMapSetAdd10Has90(b *testing.B) {
	var set = NewMapSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setAddOrHas(set, 0.1)
			}
		})
	})
}

func BenchmarkMapSetAdd50Has50(b *testing.B) {
	var set = NewMapSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setAddOrHas(set, 0.5)
			}
		})
	})
}

func BenchmarkMapSetAdd90Has10(b *testing.B) {
	var set = NewMapSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setAddOrHas(set, 0.9)
			}
		})
	})
}

func setAddOrHas(s SetInterface, addPercent float32) {
	if rand.Float32() < addPercent {
		s.Add(1)
	} else {
		s.Has(1)
	}
}
