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
}

func TestHelloDani(t *testing.T) {
	assert.Equal(t, "Hello, Dani!", HelloWorld("Dani"))
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping test")
	}
}
