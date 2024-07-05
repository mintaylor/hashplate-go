package hashplatego

import (
	"log"
	"testing"
)

func TestHelloWorldWithCN(t *testing.T) {
	input := "Hello World!"
	output := "🍢 渝G·VGUA1 🪣"

	result := Hashplate(input, defaultOptions)

	log.Println("result: ", result)

	if result != output {
		t.Errorf("HelloWorld() = %v, want %v", result, output)
	}
}

func TestHelloWorld(t *testing.T) {
	input := "Hello World!"
	output := "🦐 GP-150-UJ 🪣"

	result := Hashplate(input, Options{
		Emoji: true,
		CN:    false,
		Seed:  0,
	})

	log.Println("result: ", result)

	if result != output {
		t.Errorf("HelloWorld() = %v, want %v", result, output)
	}
}

func BenchmarkHashplateWithCN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hashplate("Hello World!", defaultOptions)
	}
}

func BenchmarkHashplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hashplate("Hello World!", Options{
			Emoji: true,
			CN:    false,
			Seed:  0,
		})
	}
}
