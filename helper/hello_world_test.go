package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	//Before Test
	fmt.Println("Before Test")
	//Run Test
	m.Run()
	//After Test
	fmt.Println("After Test")
}

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "Hello, World!", HelloWorld("World"))
	//Sub Test
	t.Run("asnur", func(t *testing.T) {
		assert.Equal(t, "Hello, asnur!", HelloWorld("asnur"))
	})
}

func TestHelloDani(t *testing.T) {
	assert.Equal(t, "Hello, Dani!", HelloWorld("Dani"))
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping test")
	}
}

//Table Test Consept
func TestTableHello(t *testing.T) {
	var tests = []struct {
		name string
		want string
	}{
		{"asnur", "Hello, asnur!"},
		{"Dani", "Hello, Dani!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, HelloWorld(tt.name))
		})
	}
}

//Benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dani")
	}
}

//Sub Bencmhmark and Table Benchmark
func BenchmarkHelloWorld_Sub(b *testing.B) {
	var bench = []struct {
		name string
	}{
		{"asnur"},
		{"Dani"},
	}
	for _, bb := range bench {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bb.name)
			}
		})
	}
}
